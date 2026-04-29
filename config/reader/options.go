package reader

import (
	"io"
	"log/slog"
	"time"

	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/config"
	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

type option interface {
	apply(*configReader)
}

type optionFunction struct {
	config.StubOption
	f func(*configReader)
}

func (f optionFunction) apply(c *configReader) {
	f.f(c)
}

func optionFunc(f func(*configReader)) config.Option {
	return &optionFunction{
		config.StubOption{},
		f,
	}
}

var _ config.Option = (*optionFunction)(nil)

func WithPreloadedConfig(cfg *config.Config) config.Option {
	return optionFunc(func(c *configReader) {
		c.config = cfg
		c.preloadedConfig = true
	})
}

func WithPath(path string) config.Option {
	return optionFunc(func(c *configReader) {
		c.path = path
	})
}

func WithProfileName(profileName string) config.Option {
	return optionFunc(func(c *configReader) {
		c.profileName = profileName
	})
}

func WithoutParentID() config.Option {
	return optionFunc(func(c *configReader) {
		c.noParentID = true
	})
}

func WithNoParentID(noParentID bool) config.Option {
	return optionFunc(func(c *configReader) {
		c.noParentID = noParentID
	})
}

func WithProfileEnv(profileEnv string) config.Option {
	return optionFunc(func(c *configReader) {
		c.profileEnv = profileEnv
	})
}

func WithTokenEnv(tokenEnv string) config.Option {
	return optionFunc(func(c *configReader) {
		c.tokenEnv = tokenEnv
	})
}

func WithEndpointEnv(endpointEnv string) config.Option {
	return optionFunc(func(c *configReader) {
		c.endpointEnv = endpointEnv
	})
}

func WithCustomEndpoint(customEndpoint string) config.Option {
	return optionFunc(func(c *configReader) {
		c.customEndpoint = customEndpoint
	})
}

func WithoutEnvVars() config.Option {
	return optionFunc(func(c *configReader) {
		c.profileEnv = ""
		c.tokenEnv = ""
	})
}

func WithFileRefreshPeriod(fileRefreshPeriod time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.fileRefreshPeriod = fileRefreshPeriod
	})
}

// file cache control
func WithoutFileCache() config.Option {
	return optionFunc(func(c *configReader) {
		c.noFileCache = true
	})
}

func WithNoFileCache(noFileCache bool) config.Option {
	return optionFunc(func(c *configReader) {
		c.noFileCache = noFileCache
	})
}

func WithCacheFileName(name string) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithFileCacheFileName(name))
	})
}

func WithClientID(clientID string) config.Option {
	return optionFunc(func(c *configReader) {
		c.clientID = clientID
	})
}

func WithLogger(logger *slog.Logger) config.Option {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	return optionFunc(func(c *configReader) {
		c.logger = logger
	})
}

// browser/open control for federation flow
func WithoutBrowserOpen() config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithFederationNoBrowserOpen())
	})
}

func WithNoBrowserOpen(noBrowserOpen bool) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithFederationNoBrowserOpenFlag(noBrowserOpen))
	})
}

// deferred client function for service-account auth
func WithDeferredClientFunc(f func() (iampb.TokenExchangeServiceClient, error)) config.Option {
	return optionFunc(func(c *configReader) {
		c.deferredClientFunc = f
	})
}

func WithWriter(w io.Writer) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithFederationWriter(w))
	})
}

// cached token tuning
func WithCachedTokenLifetimeFraction(frac float64) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithCachedTokenerLifetime(frac))
	})
}

func WithCachedTokenInitialRetryDelay(d time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithCachedTokenerInitialRetry(d))
	})
}

func WithCachedTokenRetryMultiplier(mult float64) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithCachedTokenerRetryMultiplier(mult))
	})
}

func WithCachedTokenMaxRetryDelay(d time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithCachedTokenerMaxRetry(d))
	})
}

// WithAuthOptions appends low-level auth options that will be forwarded to the
// tokeners and decorators constructed by the config reader.
//
// Each option is applied only to the auth components that support it.
func WithAuthOptions(opts ...auth.Option) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, opts...)
	})
}

func WithTokenSafetyMargin(d time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.authOptions = append(c.authOptions, auth.WithFileCacheSafetyMargin(d))
	})
}
