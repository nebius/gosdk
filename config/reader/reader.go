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

	"go.yaml.in/yaml/v3"

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

	path            string
	profileName     string
	noParentID      bool
	profileEnv      string
	tokenEnv        string
	endpointEnv     string
	customEndpoint  string
	metadataProbe   metadataProbe
	preloadedConfig bool

	// credentials options
	noFileCache        bool
	fileRefreshPeriod  time.Duration
	clientID           string
	logger             *slog.Logger
	metrics            Metrics
	deferredClientFunc func() (iampb.TokenExchangeServiceClient, error)
	authOptions        []auth.Option

	impersonateServiceAccountID string

	configParsed bool
}

func NewConfigReader(options ...config.Option) config.ConfigInterface {
	r := &configReader{
		tokenEnv:          constants.TokenEnv,
		profileEnv:        constants.ProfileEnv,
		endpointEnv:       constants.EndpointEnv,
		fileRefreshPeriod: 5 * time.Minute,
		metadataProbe:     defaultMetadataProbe(),
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
	start := time.Now()
	source := "file"
	if r.preloadedConfig {
		source = "preloaded"
	}
	retErr := r.loadProfile(ctx)

	result := metricResultSuccess
	if retErr != nil {
		result = metricResultError
	}
	configLoad(ctx, r.metrics, source, result, time.Since(start))
	return retErr
}

func (r *configReader) loadProfile(ctx context.Context) error {
	if err := r.ensureConfigLoaded(ctx); err != nil {
		return err
	}
	r.loadCustomEndpoint(ctx)
	r.loadProfileName(ctx)
	if err := r.selectProfile(ctx); err != nil {
		return err
	}
	r.logger.DebugContext(ctx, "loaded profile", slog.Any("profile", r.profile))
	r.configParsed = true
	return nil
}

func (r *configReader) ensureConfigLoaded(ctx context.Context) error {
	if r.preloadedConfig {
		r.logger.DebugContext(ctx, "config is preloaded, skipping file load")
		return nil
	}
	return r.load(ctx)
}

func (r *configReader) loadCustomEndpoint(ctx context.Context) {
	if r.customEndpoint == "" && r.endpointEnv != "" {
		r.customEndpoint = strings.TrimSpace(os.Getenv(r.endpointEnv))
		if r.customEndpoint != "" {
			r.logger.DebugContext(
				ctx,
				"will load custom endpoint from environment variable",
				slog.String("env", r.endpointEnv),
			)
		}
	}
	if r.customEndpoint != "" {
		r.logger.DebugContext(
			ctx,
			"overriding profile endpoint with custom endpoint",
			slog.String("endpoint", r.customEndpoint),
		)
	}
}

func (r *configReader) loadProfileName(ctx context.Context) {
	if r.profileName != "" || r.profileEnv == "" {
		return
	}
	r.profileName = strings.TrimSpace(os.Getenv(r.profileEnv))
	if r.profileName != "" {
		r.logger.DebugContext(ctx, "will load profile from environment variable", slog.String("env", r.profileEnv))
	}
}

func (r *configReader) selectProfile(ctx context.Context) error {
	var err error
	if r.profileName == "" {
		r.logger.DebugContext(ctx, "no profile specified, loading default profile")
		r.profile, err = r.config.GetDefaultProfile()
	} else {
		r.logger.DebugContext(ctx, "loading specified profile", slog.String("profile", r.profileName))
		r.profile, err = r.config.GetProfile(r.profileName)
	}
	if err != nil {
		return err
	}
	if r.profile == nil {
		return config.NewGetProfileError(fmt.Errorf("no profile found"), r.config.Profiles)
	}
	return nil
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
			if r.setVMProfile(ctx) {
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
	start := time.Now()
	tokener, source, retErr := r.resolveCredentials(ctx)

	result := metricResultSuccess
	if retErr != nil {
		result = metricResultError
	}
	credentialsResolve(ctx, r.metrics, source, result, time.Since(start))
	return tokener, retErr
}

func (r *configReader) resolveCredentials(ctx context.Context) (auth.BearerTokener, string, error) {
	authOpts := r.authOptionsWithLogger()
	if tokener, source, ok, err := r.resolveRawCredentials(ctx, authOpts); ok {
		return tokener, source, err
	}
	return r.resolveProfileCredentials(ctx, authOpts)
}

func (r *configReader) resolveRawCredentials(
	ctx context.Context,
	authOpts []auth.Option,
) (auth.BearerTokener, string, bool, error) {
	if r.tokenEnv != "" {
		token := strings.TrimSpace(os.Getenv(r.tokenEnv))
		if token != "" {
			r.logger.DebugContext(ctx, "loading token from environment variable", slog.String("env", r.tokenEnv))
			return auth.NewStaticTokener(token, authOpts...), "env", true, nil
		}
	}
	if r.profile.TokenFile != "" {
		tokener, err := r.newFileTokener(ctx, authOpts)
		return tokener, "token-file", true, err
	}
	if r.profile.TokenEndpoint != "" {
		tokener, err := r.newIMDSTokener(ctx, authOpts)
		return tokener, "imds", true, err
	}
	return nil, "unknown", false, nil
}

func (r *configReader) newFileTokener(ctx context.Context, authOpts []auth.Option) (auth.BearerTokener, error) {
	r.logger.DebugContext(ctx, "using token from file", slog.String("file", r.profile.TokenFile))
	tokener, err := auth.NewFileTokener(
		r.profile.TokenFile,
		r.fileRefreshPeriod,
		authOpts...,
	)
	if err != nil {
		return nil, fmt.Errorf("create token file reader %q: %w", r.profile.TokenFile, err)
	}
	return tokener, nil
}

func (r *configReader) newIMDSTokener(ctx context.Context, authOpts []auth.Option) (auth.BearerTokener, error) {
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

func (r *configReader) resolveProfileCredentials(
	ctx context.Context,
	authOpts []auth.Option,
) (auth.BearerTokener, string, error) {
	switch r.profile.AuthType {
	case config.AuthTypeServiceAccount:
		tokener, err := r.getServiceAccountCredentials(ctx)
		return tokener, "service-account", err
	case config.AuthTypeFederation:
		tokener, err := r.getFederationCredentials(ctx, authOpts)
		return tokener, "federation", err
	default:
		return nil, "unknown", config.NewError(fmt.Errorf("unsupported auth type: %s", r.profile.AuthType))
	}
}

func (r *configReader) getFederationCredentials(
	ctx context.Context,
	authOpts []auth.Option,
) (auth.BearerTokener, error) {
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
		return auth.NewTypedNameWrapper(
			fmt.Sprintf("service-account/%s/%s", saID, pkID),
			"service-account",
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

func (r *configReader) AddImpersonationIfSet(
	ctx context.Context, tokener auth.BearerTokener,
) (auth.BearerTokener, error) {
	if tokener != nil && (r.impersonateServiceAccountID != "" || (r.profile != nil &&
		r.profile.ImpersonateServiceAccountID != "")) {
		var impersonateSAID string
		if r.impersonateServiceAccountID != "" {
			impersonateSAID = r.impersonateServiceAccountID
		} else if r.profile != nil {
			impersonateSAID = r.profile.ImpersonateServiceAccountID
		}
		r.logger.DebugContext(ctx, "impersonating service account",
			slog.String("service_account_id", impersonateSAID),
		)
		// cache is necessary if the client sends multiple requests during one
		// execution for example, Operation.Wait after create/update/delete
		return auth.NewCachedTokener(
			r.newImpersonatedBearerTokener(impersonateSAID, tokener),
			r.authOptionsWithLogger()...,
		), nil
	}
	return tokener, nil
}

func (r *configReader) newImpersonatedBearerTokener(
	serviceAccountID string,
	tokener auth.BearerTokener,
) auth.BearerTokener {
	return auth.NewExchangeImpersonatedBearerTokener(
		serviceAccountID,
		tokener,
		r.deferredClientFunc,
	)
}

func (r *configReader) GetAuthType() config.AuthType {
	return r.profile.AuthType
}
