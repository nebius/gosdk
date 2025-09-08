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
	"github.com/nebius/gosdk/conn"
	"github.com/nebius/gosdk/serviceerror"
	"github.com/nebius/gosdk/services/nebius"
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
}

const (
	defaultTimeout  = 1 * time.Minute
	defaultRetries  = 3
	defaultPerRetry = 10 * time.Second
)

// New creates a new [SDK] with the provided options.
// By default, it also performs any necessary I/O operations.
// To separate I/O operations from instantiation, use the [WithExplicitInit] option.
// The context provided to [New] can be short-lived, as it is used only for the initial setup.
// SDK may span goroutines that will use the context returned by [SDK.Context].
// If you want to stop the goroutines, you should call [SDK.Close].
func New(ctx context.Context, opts ...Option) (*SDK, error) { //nolint:funlen
	domain := "api.nebius.cloud:443"

	credentials := NoCredentials()
	handler := slog.Handler(NoopHandler{})
	explicitInit := false
	timeout := defaultTimeout

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
		}
	}

	logger := slog.New(handler)
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

	if retryOpts != nil {
		retryOptsFull := []retry.CallOption{
			retry.WithMax(defaultRetries),
			retry.WithPerRetryTimeout(defaultPerRetry),
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
		grpc.WithUserAgent(userAgent),
	)

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

	// apply timeout after retry and auth interceptors for each request
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(conn.UnaryClientTimeoutInterceptor(timeout)))

	dialer := conn.NewDialer(logger, dialOpts...)
	closes = append(closes, dialer.Close)
	inits = append(inits, customInits...)

	sdkContext, cancel := context.WithCancel(context.WithoutCancel(ctx))

	sdk := &SDK{
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
		dialer:  dialer,
		tokener: tokener,
		inits:   inits,
		closes:  closes,
		ctx:     sdkContext,
		cancel:  cancel,
	}

	if !explicitInit {
		errI := sdk.Init(ctx)
		if errI != nil {
			return nil, errI
		}
	}

	return sdk, nil
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
