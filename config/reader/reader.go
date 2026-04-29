package reader

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/nebius/gosdk/auth"
	_ "github.com/nebius/gosdk/auth/federation" // register temporary federation bridge while deprecated auth/federation handles remain
	"github.com/nebius/gosdk/config"
	"github.com/nebius/gosdk/config/paths"
	"github.com/nebius/gosdk/constants"
	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

type configReader struct {
	config  *config.Config
	profile *config.Profile

	path           string
	profileName    string
	noParentID     bool
	profileEnv     string
	tokenEnv       string
	endpointEnv    string
	customEndpoint string

	preloadedConfig bool

	// credentials options
	noFileCache        bool
	fileRefreshPeriod  time.Duration
	clientID           string
	logger             *slog.Logger
	deferredClientFunc func() (iampb.TokenExchangeServiceClient, error)
	authOptions        []auth.Option

	configParsed bool
}

func NewConfigReader(options ...config.Option) config.ConfigInterface {
	r := &configReader{
		tokenEnv:          constants.TokenEnv,
		profileEnv:        constants.ProfileEnv,
		endpointEnv:       constants.EndpointEnv,
		fileRefreshPeriod: 5 * time.Minute,
	}
	r.SetOptions(options...)
	if r.logger == nil {
		r.logger = slog.New(slog.DiscardHandler)
	}
	if !r.preloadedConfig || r.config == nil {
		r.config = config.NewConfig()
		r.preloadedConfig = false
	}
	return r
}

func (r *configReader) LoadIfNeeded(ctx context.Context) error {
	if r.configParsed {
		r.logger.DebugContext(ctx, "configuration already loaded, skipping")
		return nil
	}
	return r.Load(ctx)
}

func (r *configReader) Load(ctx context.Context) error {
	if r.preloadedConfig {
		r.logger.DebugContext(ctx, "config is preloaded, skipping file load")
	} else {
		if err := r.load(ctx); err != nil {
			return err
		}
	}
	if r.customEndpoint == "" && r.endpointEnv != "" {
		r.customEndpoint = strings.TrimSpace(os.Getenv(r.endpointEnv))
		if r.customEndpoint != "" {
			r.logger.DebugContext(ctx, "will load custom endpoint from environment variable", slog.String("env", r.endpointEnv))
		}
	}
	if r.customEndpoint != "" {
		r.logger.DebugContext(ctx, "overriding profile endpoint with custom endpoint", slog.String("endpoint", r.customEndpoint))
	}
	if r.profileName == "" && r.profileEnv != "" {
		r.profileName = strings.TrimSpace(os.Getenv(r.profileEnv))
		if r.profileName != "" {
			r.logger.DebugContext(ctx, "will load profile from environment variable", slog.String("env", r.profileEnv))
		}
	}
	if r.profileName == "" {
		r.logger.DebugContext(ctx, "no profile specified, loading default profile")
		profile, err := r.config.GetDefaultProfile()
		if err != nil {
			return err
		}
		r.profile = profile
	} else {
		r.logger.DebugContext(ctx, "loading specified profile", slog.String("profile", r.profileName))
		profile, err := r.config.GetProfile(r.profileName)
		if err != nil {
			return err
		}
		r.profile = profile
	}
	if r.profile == nil {
		return config.NewGetProfileError(fmt.Errorf("no profile found"), r.config.Profiles)
	}
	r.logger.DebugContext(ctx, "loaded profile", slog.Any("profile", r.profile))
	r.configParsed = true
	return nil
}

func (r *configReader) setVMProfile() bool {
	profileName := "virtual"
	tokenFilePath := "/mnt/cloud-metadata/token"
	if stat, statErr := os.Stat(tokenFilePath); statErr == nil && !stat.IsDir() {
		r.config.Default = profileName
		r.config.Profiles[profileName] = &config.Profile{
			Name:      profileName,
			Endpoint:  paths.DefaultAPIEndpoint,
			TokenFile: tokenFilePath,
		}
		r.profile = r.config.Profiles[profileName]
		return true
	}
	return false
}
func (r *configReader) load(ctx context.Context) error {
	path, err := r.GetConfigPath()
	if err != nil {
		return fmt.Errorf("get config path: %w", err)
	}
	r.logger.DebugContext(ctx, "loading config", slog.String("path", path))
	bytes, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			if r.setVMProfile() {
				r.logger.DebugContext(ctx, "using virtual machine profile")
				return nil
			}
			return config.NewMissingConfigError(err)
		}
		return config.NewError(err)
	}
	err = yaml.Unmarshal(bytes, r.config)
	if err != nil {
		return config.NewError(err)
	}
	for name, profile := range r.config.Profiles {
		profile.Name = name
	}
	return nil
}

func (r *configReader) GetConfigPath() (string, error) {
	if r.path == "" {
		path, err := paths.GetDefaultConfigPath()
		if err != nil {
			return "", fmt.Errorf("get default config path: %w", err)
		}
		r.path = path
	}
	return r.path, nil
}

func (r *configReader) GetConfig() *config.Config {
	return r.config
}

func (r *configReader) SetOptions(options ...config.Option) {
	for _, opt := range options {
		localOpt, ok := opt.(option)
		if !ok {
			continue
		}
		localOpt.apply(r)
	}
}

func (r *configReader) CurrentProfileName() string {
	if r.profile != nil {
		return r.profile.Name
	}
	return ""
}

func (r *configReader) ParentID() string {
	if r.noParentID {
		return ""
	}
	return r.profile.ParentID
}

func (r *configReader) TenantID() string {
	if r.noParentID {
		return ""
	}
	return r.profile.TenantID
}

func (r *configReader) Endpoint() string {
	if r.customEndpoint != "" {
		return r.customEndpoint
	}
	return r.profile.Endpoint
}

func (r *configReader) authOptionsWithLogger() []auth.Option {
	opts := make([]auth.Option, 0, len(r.authOptions)+1)
	opts = append(opts, auth.WithLogger(r.logger))
	opts = append(opts, r.authOptions...)
	return opts
}

func (r *configReader) GetCredentials(ctx context.Context) (auth.BearerTokener, error) {
	authOpts := r.authOptionsWithLogger()

	// 1. Environment token
	if r.tokenEnv != "" {
		token := strings.TrimSpace(os.Getenv(r.tokenEnv))
		if token != "" {
			r.logger.DebugContext(ctx, "loading token from environment variable", slog.String("env", r.tokenEnv))
			return auth.StaticBearerToken(token), nil
		}
	}
	// 2. Raw token file
	if r.profile.TokenFile != "" {
		r.logger.DebugContext(ctx, "using token from file", slog.String("file", r.profile.TokenFile))
		info, err := os.Stat(r.profile.TokenFile)
		if err != nil {
			return nil, config.NewError(fmt.Errorf("stat token file %q: %w", r.profile.TokenFile, err))
		}
		if info.IsDir() {
			return nil, config.NewError(fmt.Errorf("token file %q is a directory", r.profile.TokenFile))
		}
		tokener, err := auth.NewFileTokener(r.profile.TokenFile, r.fileRefreshPeriod)
		if err != nil {
			return nil, fmt.Errorf("create token file reader %q: %w", r.profile.TokenFile, err)
		}
		return tokener, nil
	}

	// 3. Raw token endpoint
	if r.profile.TokenEndpoint != "" {
		r.logger.DebugContext(ctx, "using token from imds", slog.String("endpoint", r.profile.TokenEndpoint))
		tokener, err := auth.NewIMDSTokenizer(
			r.profile.TokenEndpoint,
			authOpts...,
		)
		if err != nil {
			return nil, fmt.Errorf("create token endpoint reader %q: %w", r.profile.TokenEndpoint, err)
		}
		return tokener, nil
	}

	switch r.profile.AuthType {
	case config.AuthTypeServiceAccount:
		return r.getServiceAccountCredentials(ctx)
	case config.AuthTypeFederation:
		r.logger.DebugContext(ctx, "using federation authentication",
			slog.String("client_id", r.clientID),
			slog.String("federation_endpoint", r.profile.FederationEndpoint),
			slog.String("federation_id", r.profile.FederationID),
			slog.String("profile", r.profile.Name),
		)
		if r.clientID == "" {
			return nil, fmt.Errorf("missing client ID for federation auth")
		}
		tokener := auth.NewFederationTokener(
			r.clientID,
			r.profile.FederationEndpoint,
			r.profile.FederationID,
			r.profile.Name,
			authOpts...,
		)
		r.logger.DebugContext(ctx, "using file-cached federation tokener")
		return auth.NewInAppSyncTokener(auth.NewFileCacheTokener(tokener, authOpts...)), nil
	default:
		return nil, config.NewError(fmt.Errorf("unsupported auth type: %s", r.profile.AuthType))
	}
}

// getServiceAccountCredentials handles all service account auth permutations.
func (r *configReader) getServiceAccountCredentials(ctx context.Context) (auth.BearerTokener, error) {
	r.logger.DebugContext(ctx, "using service account authentication",
		slog.String("service_account_id", r.profile.ServiceAccountID),
		slog.String("public_key_id", r.profile.PublicKeyID),
		slog.String("private_key_file_path", r.profile.PrivateKeyFilePath),
		slog.String("service_account_credentials_file_path", r.profile.ServiceAccountCredentialsFilePath),
		slog.String("federated_subject_credentials_file_path", r.profile.FederatedSubjectCredentialsFilePath),
	)
	if r.deferredClientFunc == nil {
		return nil, errors.New("deferred client function is not set")
	}
	resolvedTokener, err := r.resolveServiceAccountTokener(ctx)
	if err != nil {
		return nil, err
	}
	if r.profile.FederatedSubjectCredentialsFilePath != "" && resolvedTokener != nil && !strings.HasPrefix(resolvedTokener.Name(), "federated-credentials/") {
		r.logger.WarnContext(ctx, "federated subject credentials file is ignored because key-based credentials are used")
	}
	authOpts := r.authOptionsWithLogger()
	if r.noFileCache {
		r.logger.DebugContext(ctx, "using in-memory cached service account tokener")
		return auth.NewCachedTokener(
			resolvedTokener,
			authOpts...,
		), nil
	}
	r.logger.DebugContext(ctx, "using file-cached service account tokener")
	return auth.NewFileCacheTokener(resolvedTokener, authOpts...), nil
}

func (r *configReader) resolveServiceAccountTokener(ctx context.Context) (auth.NamedTokener, error) {
	wrap := func(reader auth.ServiceAccountReader, saID, pkID string) auth.NamedTokener {
		return auth.NewNameWrapper(
			fmt.Sprintf("service-account/%s/%s", saID, pkID),
			auth.NewExchangeableBearerTokenerWithDeferredClient(
				auth.NewServiceAccountExchangeTokenRequester(
					auth.NewCachedServiceAccount(reader),
				), r.deferredClientFunc,
			),
		)
	}

	if r.profile.FederatedSubjectCredentialsFilePath != "" &&
		r.profile.ServiceAccountID != "" &&
		r.profile.PublicKeyID == "" &&
		r.profile.PrivateKey == "" &&
		r.profile.PrivateKeyFilePath == "" &&
		r.profile.ServiceAccountCredentialsFilePath == "" {
		return auth.NewFederatedCredentialsTokener(
			r.profile.ServiceAccountID,
			r.profile.FederatedSubjectCredentialsFilePath,
			r.deferredClientFunc,
		), nil
	}

	if r.profile.ServiceAccountCredentialsFilePath != "" {
		if r.profile.ServiceAccountID != "" || r.profile.PublicKeyID != "" {
			return nil, fmt.Errorf("incomplete service account configuration: provide either (service-account-id, federated-subject-credentials-file-path) OR (service-account-credentials-file-path) OR (service-account-id, public-key-id and one of private-key / private-key-file-path)")
		}
		if r.profile.PrivateKey != "" || r.profile.PrivateKeyFilePath != "" {
			r.logger.WarnContext(ctx, "ignoring inline parameters or file private key because service account credentials file is provided")
		}
		parser := auth.NewServiceAccountCredentialsFileParser(nil, r.profile.ServiceAccountCredentialsFilePath)
		creds, err := parser.SubjectCredentials()
		if err != nil {
			return nil, fmt.Errorf("parse service account credentials file: %w", err)
		}
		return wrap(parser, creds.Subject, creds.KeyID), nil
	}

	if r.profile.PrivateKey != "" &&
		r.profile.PublicKeyID != "" &&
		r.profile.ServiceAccountID != "" {
		if r.profile.PrivateKeyFilePath != "" {
			r.logger.WarnContext(ctx, "ignoring file private key because inline private key is provided")
		}
		return wrap(
			auth.NewPrivateKeyParser([]byte(r.profile.PrivateKey), r.profile.PublicKeyID, r.profile.ServiceAccountID),
			r.profile.ServiceAccountID,
			r.profile.PublicKeyID,
		), nil
	}

	if r.profile.PrivateKeyFilePath != "" &&
		r.profile.PublicKeyID != "" &&
		r.profile.ServiceAccountID != "" {
		return wrap(
			auth.NewPrivateKeyFileParser(nil, r.profile.PrivateKeyFilePath, r.profile.PublicKeyID, r.profile.ServiceAccountID),
			r.profile.ServiceAccountID,
			r.profile.PublicKeyID,
		), nil
	}

	return nil, fmt.Errorf("incomplete service account configuration: provide either (service-account-id, federated-subject-credentials-file-path) OR (service-account-credentials-file-path) OR (service-account-id, public-key-id and one of private-key / private-key-file-path)")
}

func (r *configReader) GetAuthType() config.AuthType {
	return r.profile.AuthType
}
