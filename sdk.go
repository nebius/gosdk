package gosdk

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	grpc_creds "google.golang.org/grpc/credentials"

	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/conn"
	"github.com/nebius/gosdk/serviceerror"
	"github.com/nebius/gosdk/services/nebius"
)

// SDK is the Nebius API client.
type SDK struct {
	resolver conn.Resolver
	dialer   conn.Dialer
	tokener  auth.BearerTokener
	inits    []func(context.Context, *SDK) error
	closes   []func() error
}

// New creates a new [SDK] with the provided options.
// By default, it also performs any necessary I/O operations.
// To separate I/O operations from instantiation, use the [WithExplicitInit] option.
//
// Important: Ensure that the provided ctx is not closed before calling [SDK.Close].
func New(ctx context.Context, opts ...Option) (*SDK, error) { //nolint:funlen
	domain := "api.eu.nebius.cloud:443"

	credentials := NoCredentials()
	handler := slog.Handler(NoopHandler{})
	explicitInit := false

	customSubstitutions := make(map[string]string)
	var logOpts []logging.Option
	var customDialOpts []grpc.DialOption
	var customResolvers []conn.Resolver
	var customInits []func(context.Context, *SDK) error
	for _, opt := range opts {
		switch o := opt.(type) {
		case optionCredentials:
			credentials = o.creds
		case optionLogger:
			handler = o.handler
		case optionLoggingOptions:
			logOpts = append(logOpts, o...)
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
		}
	}

	logger := slog.New(handler)
	substitutions := map[string]string{
		"{domain}": domain,
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
			inits = append(inits, initCache(cache))
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
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(conn.UnaryClientTimeoutInterceptor(1*time.Minute)))

	dialer := conn.NewDialer(logger, dialOpts...)
	closes = append(closes, dialer.Close)
	inits = append(inits, customInits...)

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
	}

	if !explicitInit {
		errI := sdk.Init(ctx)
		if errI != nil {
			return nil, errI
		}
	}

	return sdk, nil
}

// Init finalizes the creation of the [SDK] by performing all required I/O operations.
// It is automatically called by the [New] constructor by default.
// This method should only be called manually if the [WithExplicitInit] option is used.
//
// Important: Ensure that the provided ctx is not closed before calling [SDK.Close].
func (s *SDK) Init(ctx context.Context) error {
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

// Close closes connections and leans up resources.
func (s *SDK) Close() error {
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

func initCache(cache *auth.CachedServiceTokener) func(context.Context, *SDK) error {
	const wait = 5 * time.Second
	return func(ctx context.Context, sdk *SDK) error {
		ctx, cancel := context.WithCancel(ctx)
		errs := make(chan error, 1)
		stopped := atomic.Bool{}

		stopCacheAndWaitResult := func() error {
			if !stopped.CompareAndSwap(false, true) {
				return nil
			}

			cancel()

			select {
			case err := <-errs:
				if err == context.Canceled {
					return nil
				}
				return err
			case <-time.After(wait):
				return errors.New("timed out waiting for cache")
			}
		}

		// Need to stop cache before closing dialer
		sdk.closes = append([]func() error{stopCacheAndWaitResult}, sdk.closes...)

		go func() {
			errs <- cache.Run(ctx)
		}()

		return nil
	}
}
