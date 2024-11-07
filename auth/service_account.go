package auth

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
)

type ServiceAccount struct {
	PrivateKey       *rsa.PrivateKey
	PublicKeyID      string
	ServiceAccountID string
}

type ServiceAccountReader interface {
	ServiceAccount(context.Context) (ServiceAccount, error)
}

type CachedServiceAccount struct {
	reader ServiceAccountReader
	group  singleflight.Group

	mu    sync.RWMutex
	cache *ServiceAccount
}

var _ ServiceAccountReader = (*CachedServiceAccount)(nil)

func NewCachedServiceAccount(reader ServiceAccountReader) *CachedServiceAccount {
	return &CachedServiceAccount{
		reader: reader,
		group:  singleflight.Group{},
		mu:     sync.RWMutex{},
		cache:  nil,
	}
}

func (c *CachedServiceAccount) ServiceAccount(ctx context.Context) (ServiceAccount, error) {
	c.mu.RLock()
	cache := c.cache
	c.mu.RUnlock()

	if cache != nil {
		return *cache, nil
	}

	res, err, _ := c.group.Do("", func() (interface{}, error) {
		account, err := c.reader.ServiceAccount(ctx)
		if err != nil {
			return nil, err
		}

		c.mu.Lock()
		c.cache = &account
		c.mu.Unlock()

		return account, nil
	})
	if err != nil {
		return ServiceAccount{}, err
	}

	return res.(ServiceAccount), nil
}

type StaticServiceAccount ServiceAccount

var _ ServiceAccountReader = StaticServiceAccount{}

func (s StaticServiceAccount) ServiceAccount(context.Context) (ServiceAccount, error) {
	return ServiceAccount(s), nil
}

// PrivateKeyParser is a [ServiceAccountReader] that parses a PEM encoded PKCS1 or PKCS8 private key.
type PrivateKeyParser struct {
	privateKey       []byte
	publicKeyID      string
	serviceAccountID string
}

var _ ServiceAccountReader = PrivateKeyParser{}

func NewPrivateKeyParser(privateKey []byte, publicKeyID string, serviceAccountID string) PrivateKeyParser {
	return PrivateKeyParser{
		privateKey:       privateKey,
		publicKeyID:      publicKeyID,
		serviceAccountID: serviceAccountID,
	}
}

func (p PrivateKeyParser) ServiceAccount(context.Context) (ServiceAccount, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(p.privateKey)
	if err != nil {
		return ServiceAccount{}, err
	}
	return ServiceAccount{
		PrivateKey:       privateKey,
		PublicKeyID:      p.publicKeyID,
		ServiceAccountID: p.serviceAccountID,
	}, nil
}

// PrivateKeyFileParser is a [ServiceAccountReader] that reads a file
// and parses a PEM encoded PKCS1 or PKCS8 private key.
// You can set fs = nil to use default fs.
type PrivateKeyFileParser struct {
	fs               fs.FS
	privateKeyPath   string
	publicKeyID      string
	serviceAccountID string
}

var _ ServiceAccountReader = PrivateKeyFileParser{}

func NewPrivateKeyFileParser(
	fs fs.FS,
	privateKeyPath string,
	publicKeyID string,
	serviceAccountID string,
) PrivateKeyFileParser {
	return PrivateKeyFileParser{
		fs:               fs,
		privateKeyPath:   privateKeyPath,
		publicKeyID:      publicKeyID,
		serviceAccountID: serviceAccountID,
	}
}

const homeShortcut = "~" + string(filepath.Separator)

func sanitizePath(path string) (string, error) {
	if path == "~" {
		return "", errors.New("invalid path '~'")
	}
	if strings.HasPrefix(path, homeShortcut) {
		homePath, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(
			homePath, strings.TrimPrefix(path, homeShortcut),
		)
	}
	realPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return realPath, nil
}

func fsForGlobalPath(path string) (fs.FS, string, error) {
	realPath, err := sanitizePath(path)
	if err != nil {
		return nil, "", err
	}
	dirPath, fileName := filepath.Split(realPath)
	if dirPath == "" {
		dirPath = "."
	}
	return os.DirFS(dirPath), fileName, nil
}

func (p PrivateKeyFileParser) ServiceAccount(context.Context) (ServiceAccount, error) {
	if p.fs == nil {
		var err error
		p.fs, p.privateKeyPath, err = fsForGlobalPath(p.privateKeyPath)
		if err != nil {
			return ServiceAccount{}, err
		}
	}
	encoded, err := fs.ReadFile(p.fs, p.privateKeyPath)
	if err != nil {
		return ServiceAccount{}, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(encoded)
	if err != nil {
		return ServiceAccount{}, err
	}
	return ServiceAccount{
		PrivateKey:       privateKey,
		PublicKeyID:      p.publicKeyID,
		ServiceAccountID: p.serviceAccountID,
	}, nil
}

// ServiceAccountCredentialsFileParser is a [ServiceAccountReader] that reads a json file with service account id,
// public key id and PEM encoded PKCS8 private key.
// Docs for service account credentials file format:
// https://docs.nebius.dev/en/iam/for-services/step-by-step/authentication/service-account/sa-credentials-file
//
// You can set fs = nil to use default fs.
type ServiceAccountCredentialsFileParser struct {
	fs              fs.FS
	credentialsPath string
}

var _ ServiceAccountReader = ServiceAccountCredentialsFileParser{}

func NewServiceAccountCredentialsFileParser(fs fs.FS, credentialsPath string) ServiceAccountCredentialsFileParser {
	return ServiceAccountCredentialsFileParser{
		fs:              fs,
		credentialsPath: credentialsPath,
	}
}

// Note: json field names are standard abbreviations (see, JWT standards or RFC 7515 for example)
// Fields are interpreted according to IAM doc mentioned above.
type ServiceAccountCredentials struct {
	SubjectCredentials SubjectCredentials `json:"subject-credentials"`
}

type SubjectCredentials struct {
	Type       string `json:"type,omitempty"`
	Alg        string `json:"alg"`
	PrivateKey string `json:"private-key"` // PEM encoded PKCS8 string
	KeyID      string `json:"kid"`
	Issuer     string `json:"iss"`
	Subject    string `json:"sub"`
}

func (c SubjectCredentials) validate() error {
	if (c.Type != "") && (c.Type != "JWT") {
		return fmt.Errorf("invalid service account credentials type: '%s', the only supported value is 'JWT'", c.Type)
	}

	if c.Alg != "RS256" {
		return fmt.Errorf("invalid service account algorithm: %s. Only RS256 is supported", c.Alg)
	}

	if c.Issuer != c.Subject {
		return fmt.Errorf("invalid service account subject must be the same as issuer: %s != %s", c.Issuer, c.Subject)
	}

	return nil
}

func (p ServiceAccountCredentialsFileParser) ServiceAccount(context.Context) (ServiceAccount, error) {
	creds, err := p.SubjectCredentials()
	if err != nil {
		return ServiceAccount{}, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(creds.PrivateKey))
	if err != nil {
		return ServiceAccount{}, fmt.Errorf("invalid service account private key field: %w", err)
	}
	return ServiceAccount{
		PrivateKey:       privateKey,
		PublicKeyID:      creds.KeyID,
		ServiceAccountID: creds.Subject,
	}, nil
}

func (p ServiceAccountCredentialsFileParser) SubjectCredentials() (SubjectCredentials, error) {
	if p.fs == nil {
		var err error
		p.fs, p.credentialsPath, err = fsForGlobalPath(p.credentialsPath)
		if err != nil {
			return SubjectCredentials{}, err
		}
	}
	data, err := fs.ReadFile(p.fs, p.credentialsPath)
	if err != nil {
		return SubjectCredentials{}, fmt.Errorf("can't read service account credentials file: %w", err)
	}

	var credsWithWrapper ServiceAccountCredentials
	err = json.Unmarshal(data, &credsWithWrapper)
	if err != nil {
		return SubjectCredentials{}, fmt.Errorf("parse service account creds: %w", err)
	}

	err = credsWithWrapper.SubjectCredentials.validate()
	if err != nil {
		return SubjectCredentials{}, err
	}

	return credsWithWrapper.SubjectCredentials, nil
}
