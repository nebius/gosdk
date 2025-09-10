package config

import (
	"errors"
	"fmt"
	"log/slog"
)

type AuthType string

const (
	AuthTypeFederation     AuthType = "federation"
	AuthTypeServiceAccount AuthType = "service account"
)

type Config struct {
	Default  string         `yaml:"default,omitempty"`
	Profiles ProfilesConfig `yaml:"profiles"`
}

type ProfilesConfig map[string]*Profile

type Profile struct {
	Name string `yaml:"-"`

	Endpoint           string   `yaml:"endpoint,omitempty"`
	TokenFile          string   `yaml:"token-file,omitempty"`
	AuthType           AuthType `yaml:"auth-type,omitempty"`
	FederationEndpoint string   `yaml:"federation-endpoint,omitempty"`
	FederationID       string   `yaml:"federation-id,omitempty"`
	// TODO: move the following 4 values to a nested struct (config command doesn't support that)
	ServiceAccountID string `yaml:"service-account-id,omitempty"`
	PublicKeyID      string `yaml:"public-key-id,omitempty"`
	PrivateKey       string `yaml:"private-key,omitempty"`
	// PrivateKeyFilePath points to a file containing a PEM encoded PKCS1 or PKCS8 private key.
	PrivateKeyFilePath string `yaml:"private-key-file-path,omitempty"`
	// ServiceAccountCredentialsFilePath points to a JSON credentials file that already
	// contains service account ID, public key ID and private key (see auth.ServiceAccountCredentialsFileParser format).
	ServiceAccountCredentialsFilePath string `yaml:"service-account-credentials-file-path,omitempty"`
	// Federated credentials log-in path
	FederatedSubjectCredentialsFilePath string `yaml:"federated-subject-credentials-file-path,omitempty"`

	ParentID string `yaml:"parent-id,omitempty"`
}

func (p *Profile) LogValue() slog.Value {
	return slog.AnyValue(map[string]any{
		"name":                p.Name,
		"endpoint":            p.Endpoint,
		"federation-endpoint": p.FederationEndpoint,
		"federation-id":       p.FederationID,
		"service-account-id":  p.ServiceAccountID,
		"public-key-id":       p.PublicKeyID,
		// Do not log private key
		// "private-key":         p.PrivateKey,
		"private-key-file-path":                   p.PrivateKeyFilePath,
		"service-account-credentials-file-path":   p.ServiceAccountCredentialsFilePath,
		"federated-subject-credentials-file-path": p.FederatedSubjectCredentialsFilePath,
		"auth-type":  p.AuthType,
		"parent-id":  p.ParentID,
		"token-file": p.TokenFile,
	})
}

func NewConfig() *Config {
	cfg := &Config{
		Profiles: make(ProfilesConfig),
	}
	return cfg
}

func (c *Config) GetDefaultProfile() (*Profile, error) {
	if c.Default == "" {
		switch len(c.Profiles) {
		case 0:
			return nil, NewError(NewGetProfileError(errors.New("no profile configured"), c.Profiles))
		case 1:
			for _, profile := range c.Profiles {
				return profile, nil
			}
		default:
			return nil, NewError(NewGetProfileError(
				errors.New("malformed configuration: default profile is empty"),
				c.Profiles,
			))
		}
	}
	profile, err := c.GetProfile(c.Default)
	if err != nil {
		return nil, NewError(err)
	}
	return profile, nil
}

func (c *Config) HasProfile(name string) bool {
	_, exists := c.Profiles[name]
	return exists
}

func (c *Config) GetProfile(name string) (*Profile, error) {
	profile, exists := c.Profiles[name]
	if !exists {
		return nil, NewGetProfileError(fmt.Errorf("unknown profile %q", name), c.Profiles)
	}
	return profile, nil
}
