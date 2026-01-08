package gosdk

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	grpc_creds "google.golang.org/grpc/credentials"

	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/config"
	"github.com/nebius/gosdk/config/paths"
	"github.com/nebius/gosdk/config/reader"
	"github.com/nebius/gosdk/conn"
	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
	"github.com/nebius/gosdk/serviceerror"
	"github.com/nebius/gosdk/services/nebius"
	iam "github.com/nebius/gosdk/services/nebius/iam/v1"
)

// SDK is the Nebius API client.
type SDK struct {
	ctx      context.Context
	cancel   context.CancelFunc
	resolver conn.Resolver
	dialer   conn.Dialer
	tokener  auth.BearerTokener
	inits    []func(context.Context, *SDK) error
	closes   []func() error
	isClosed atomic.Bool
	parentID string
	tenantID string
	logger   *slog.Logger
}

const (
	DefaultTimeout     = 1 * time.Minute
	DefaultRetries     = 3
	DefaultPerRetry    = DefaultTimeout / DefaultRetries
	DefaultAuthTimeout = 15 * time.Minute // should be large enough to authenticate manually
)

// New creates a new [SDK] with the provided options.
// By default, it also performs any necessary I/O operations.
// To separate I/O operations from instantiation, use the [WithExplicitInit] option.
// The context provided to [New] can be short-lived, as it is used only for the initial setup.
// SDK may span goroutines that will use the context returned by [SDK.Context].
// If you want to stop the goroutines, you should call [SDK.Close].
// Add a [WithConfigReader] option to read configuration from the standard Nebius CLI config file.
// We also recommend setting a [WithUserAgentPrefix] option to identify your application.
func New(ctx context.Context, opts ...Option) (*SDK, error) { //nolint:funlen
	var domain string
	var sdk *SDK

	var credentials Credentials = nil
	handler := slog.Handler(NoopHandler{})
	explicitInit := false
	timeout := DefaultTimeout
	authTimeout := DefaultAuthTimeout

	userAgent := "nebius-gosdk"
	goVer := runtime.Version()
	if goVer != "" {
		goVer = goVer[2:]
		if i := strings.Index(goVer, " X:"); i >= 0 {
			goVer = goVer[:i]
		}
	}
	userAgent += " ("
	userAgent += runtime.GOOS +
		" " + runtime.GOARCH + "; go/" + goVer
	userAgent += ")"

	customSubstitutions := make(map[string]string)
	var logOpts []logging.Option
	var customDialOpts []grpc.DialOption
	var customResolvers []conn.Resolver
	var customInits []func(context.Context, *SDK) error
	var configReader config.ConfigInterface = nil
	var parentID string
	var noParentID bool
	var tenantID string
	retryOpts := []retry.CallOption{}
	for _, opt := range opts {
		switch o := opt.(type) {
		case optionCredentials:
			credentials = o.creds
		case optionLogger:
			handler = o.handler
		case optionLoggingOptions:
			logOpts = append(logOpts, o...)
		case optionUserAgentPrefix:
			userAgent = string(o) + " " + userAgent
		case optionDialOpts:
			customDialOpts = append(customDialOpts, o...)
		case optionResolvers:
			customResolvers = append(customResolvers, o...)
		case optionDomain:
			if o == "" {
				return nil, errors.New("empty domain provided")
			}
			domain = string(o)
		case optionAddressTemplate:
			customSubstitutions[o.find] = o.replace
		case optionExplicitInit:
			explicitInit = bool(o)
		case optionInit:
			customInits = append(customInits, o)
		case optionRetryOptions:
			retryOpts = append(retryOpts, o...)
		case optionTimeout:
			switch {
			case o == 0:
				return nil, errors.New("zero timeout provided")
			case o < 0:
				return nil, errors.New("negative timeout provided")
			default:
				timeout = time.Duration(o)
			}
		case optionAuthTimeout:
			switch {
			case o == 0:
				return nil, errors.New("zero auth timeout provided")
			case o < 0:
				return nil, errors.New("negative auth timeout provided")
			default:
				authTimeout = time.Duration(o)
			}
		case optionConfigReader:
			configReader = o.configReader
		case optionParentID:
			parentID = string(o)
			noParentID = false // to ensure that WithParentID overrides the previous WithNoParentID
		case optionNoParentID:
			noParentID = bool(o)
		case optionTenantID:
			tenantID = string(o)
		}
	}
	logger := slog.New(handler)
	if configReader != nil {
		logger.DebugContext(ctx, "SDK is initialized with config reader")
		configReader.SetOptions(
			reader.WithLogger(logger),
			reader.WithDeferredClientFunc(func() (iampb.TokenExchangeServiceClient, error) {
				if sdk == nil {
					return nil, errors.New("SDK is not initialized")
				}
				return iam.NewTokenExchangeService(sdk), nil
			}),
		)
		if err := configReader.LoadIfNeeded(ctx); err != nil {
			return nil, fmt.Errorf("load config: %w", err)
		}

		if domain == "" {
			if configReader.Endpoint() == "" {
				domain = paths.DefaultAPIEndpoint
				logger.WarnContext(ctx, "missing profile's endpoint, using default",
					slog.String("endpoint", domain),
					slog.String("profile", configReader.CurrentProfileName()),
				)
			} else {
				logger.DebugContext(ctx, "using endpoint from config reader",
					slog.String("endpoint", configReader.Endpoint()),
				)
				domain = configReader.Endpoint()
			}
		}
		if credentials == nil {
			tokener, err := configReader.GetCredentials(ctx)
			if err != nil {
				return nil, fmt.Errorf("get credentials: %w", err)
			}
			credentials = credsTokener{
				tokener: tokener,
			}
		} else {
			logger.DebugContext(ctx, "credentials provided by options, ignoring credentials from config reader",
				slog.Any("credentials", credentials),
			)
		}
		if parentID == "" {
			parentID = configReader.ParentID()
		}
		if tenantID == "" {
			tenantID = configReader.TenantID()
		}

	}
	if credentials == nil {
		credentials = NoCredentials()
	}

	if domain == "" {
		domain = paths.DefaultAPIEndpoint
	}

	substitutions := map[string]string{
		"{domain}": ensurePort(domain),
	}
	for find, replace := range customSubstitutions {
		substitutions[find] = replace
	}

	var inits []func(context.Context, *SDK) error
	var closes []func() error
	var tokener auth.BearerTokener

	explicitSelector := false
	credsMap := map[auth.Selector]Credentials{
		auth.Select("default"): credentials,
	}
	switch c := credentials.(type) { //nolint:gocritic // complains about single case in public gosdk
	case credsOneOf:
		if len(c) == 0 {
			return nil, errors.New("credentials map is empty")
		}
		explicitSelector = true
		credsMap = c
	}

	auths := make(map[auth.Selector]auth.Authenticator, len(credsMap))
	for selector, creds := range credsMap {
		switch c := creds.(type) {
		case credsNoCreds:
			auths[selector] = nil
		case credsAuthenticator:
			auths[selector] = c.auth
		case credsTokener:
			tokener = c.tokener
			auths[selector] = auth.NewAuthenticatorFromBearerTokener(tokener)
		case credsServiceAccount:
			t := auth.NewExchangeableBearerTokener(auth.NewServiceAccountExchangeTokenRequester(c.reader), nil)
			inits = append(inits, func(_ context.Context, sdk *SDK) error {
				t.SetClient(sdk.Services().IAM().V1().TokenExchange())
				return nil
			})

			//nolint:mnd // 90%, 1s and 1m are recommended values by IAM team
			cache := auth.NewCachedServiceTokener(logger, t, 0.9, 1*time.Second, 1.5, 1*time.Minute)
			inits = append(inits, initCache(cache, logger))
			tokener = cache

			auths[selector] = auth.NewAuthenticatorFromBearerTokener(tokener)
		case credsPropagate:
			auths[selector] = auth.NewPropagateAuthorizationHeader()
		case credsOneOf:
			return nil, errors.New("one of credentials can not be nested")
		default:
			return nil, fmt.Errorf("unknown credentials type: %T", creds)
		}
	}

	dialOpts := []grpc.DialOption{}

	dialOpts = append(dialOpts,
		grpc.WithChainUnaryInterceptor(conn.IdempotencyKeyInterceptor),
		grpc.WithTransportCredentials(grpc_creds.NewTLS(nil)), // tls enabled by default
	)

	dialOpts = append(dialOpts, customDialOpts...)

	dialOpts = append(dialOpts, grpc.WithUserAgent(userAgent))

	// apply auth timeout, it must cover both authentication and request
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(conn.UnaryClientTimeoutInterceptor(authTimeout)))

	// do not add interceptors if no credentials provided
	for _, a := range auths {
		if a != nil {
			dialOpts = append(dialOpts,
				grpc.WithChainUnaryInterceptor(auth.UnaryClientInterceptor(logger, auths, explicitSelector)),
				grpc.WithChainStreamInterceptor(auth.StreamClientInterceptor(logger, auths, explicitSelector)),
			)
			break
		}
	}

	// apply overall timeout
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(conn.UnaryClientTimeoutInterceptor(timeout)))

	if retryOpts != nil {
		retryOptsFull := []retry.CallOption{
			retry.WithMax(DefaultRetries),
			retry.WithPerRetryTimeout(DefaultPerRetry),
			conn.WithServiceRetry(),
		}
		retryOptsFull = append(retryOptsFull, retryOpts...)
		dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(
			retry.UnaryClientInterceptor(
				retryOptsFull...,
			),
		))
		dialOpts = append(dialOpts, grpc.WithChainStreamInterceptor(
			retry.StreamClientInterceptor(
				retryOptsFull...,
			),
		))
	}

	dialOpts = append(dialOpts,
		grpc.WithChainUnaryInterceptor(
			unaryLoggingInterceptor(logger, logOpts...),
			serviceerror.UnaryClientInterceptor,
		),
		grpc.WithChainStreamInterceptor(
			streamLoggingInterceptor(logger, logOpts...),
			serviceerror.StreamClientInterceptor,
		),
	)

	dialer := conn.NewDialer(logger, dialOpts...)
	closes = append(closes, dialer.Close)
	inits = append(inits, customInits...)

	sdkContext, cancel := context.WithCancel(context.WithoutCancel(ctx))

	if noParentID {
		parentID = ""
		tenantID = ""
		logger.DebugContext(ctx, "parent and tenant ID are disabled by options")
	}

	sdk = &SDK{
		resolver: conn.NewContextResolver(
			logger,
			conn.NewCachedResolver(
				logger,
				conn.NewTemplateExpander(
					substitutions,
					conn.NewResolversChain(append(customResolvers, conn.NewConventionResolver())...),
				),
			),
		),
		dialer:   dialer,
		tokener:  tokener,
		inits:    inits,
		closes:   closes,
		ctx:      sdkContext,
		cancel:   cancel,
		parentID: parentID,
		tenantID: tenantID,
		logger:   logger,
	}

	if !explicitInit {
		errI := sdk.Init(ctx)
		if errI != nil {
			return nil, errI
		}
	}

	return sdk, nil
}

func (s *SDK) GetLogger() *slog.Logger {
	return s.logger
}

// Context returns a long-lived context for the internal SDK operations
// that lives as long as the SDK itself and is canceled when the SDK is closed
// (by calling [SDK.Close]).
func (s *SDK) Context() context.Context {
	return s.ctx
}

// Init finalizes the creation of the [SDK] by performing all required I/O operations.
// It is automatically called by the [New] constructor by default.
// This method should only be called manually if the [WithExplicitInit] option is used.
func (s *SDK) Init(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, init := range s.inits {
		err := init(ctx, s)
		if err != nil {
			return err
		}
	}
	s.inits = nil // reset to prevent issues if Init() is called multiple times or called without [WithExplicitInit]
	return nil
}

// ParentID returns the parent ID saved in the SDK, if any.
func (s *SDK) ParentID() string {
	return s.parentID
}

// TenantID returns the tenant ID saved in the SDK, if any.
func (s *SDK) TenantID() string {
	return s.tenantID
}

// Services is a fluent interface to get a Nebius service client.
func (s *SDK) Services() nebius.Services {
	return nebius.New(s)
}

// Close closes connections and cleans up resources.
func (s *SDK) Close() error {
	if s.isClosed.Swap(true) {
		return nil
	}
	defer s.cancel()

	var errs error
	for _, c := range s.closes {
		err := c()
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return errs
}

// BearerToken returns the token used to authorize gRPC requests.
func (s *SDK) BearerToken(ctx context.Context) (auth.BearerToken, error) {
	if s.tokener == nil {
		return auth.BearerToken{}, errors.New("tokener not initialized")
	}
	return s.tokener.BearerToken(ctx)
}

// Resolve returns an address of a specific service.
func (s *SDK) Resolve(ctx context.Context, id conn.ServiceID) (conn.Address, error) {
	return s.resolver.Resolve(ctx, id)
}

// Dial returns an underlying gRPC client connection for a specific address.
func (s *SDK) Dial(ctx context.Context, address conn.Address) (grpc.ClientConnInterface, error) {
	return s.dialer.Dial(ctx, address)
}

func initCache(cache *auth.CachedServiceTokener, logger *slog.Logger) func(context.Context, *SDK) error {
	return func(_ context.Context, sdk *SDK) error {
		go func() {
			err := cache.Run(sdk.Context())
			if err != nil && !errors.Is(err, context.Canceled) {
				logger.ErrorContext(
					sdk.Context(),
					"background token refresh failed",
					slog.Any("error", err),
				)
			}
		}()

		return nil
	}
}

func ensurePort(domain string) string {
	if !strings.Contains(domain, ":") {
		return domain + ":443"
	}
	return domain
}
