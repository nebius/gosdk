package reader

import (
	"io"
	"log/slog"
	"time"

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
		c.cacheFileName = name
	})
}

func WithClientID(clientID string) config.Option {
	return optionFunc(func(c *configReader) {
		c.clientID = clientID
	})
}

func WithLogger(logger *slog.Logger) config.Option {
	return optionFunc(func(c *configReader) {
		c.logger = logger
	})
}

// browser/open control for federation flow
func WithoutBrowserOpen() config.Option {
	return optionFunc(func(c *configReader) {
		c.noBrowserOpen = true
	})
}

func WithNoBrowserOpen(noBrowserOpen bool) config.Option {
	return optionFunc(func(c *configReader) {
		c.noBrowserOpen = noBrowserOpen
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
		c.writer = w
	})
}

// cached token tuning
func WithCachedTokenLifetimeFraction(frac float64) config.Option {
	return optionFunc(func(c *configReader) {
		c.cachedTokenLifetimeFraction = frac
	})
}

func WithCachedTokenInitialRetryDelay(d time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.cachedTokenInitialRetryDelay = d
	})
}

func WithCachedTokenRetryMultiplier(mult float64) config.Option {
	return optionFunc(func(c *configReader) {
		c.cachedTokenRetryMultiplier = mult
	})
}

func WithCachedTokenMaxRetryDelay(d time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.cachedTokenMaxRetryDelay = d
	})
}

func WithTokenSafetyMargin(d time.Duration) config.Option {
	return optionFunc(func(c *configReader) {
		c.tokenSafetyMargin = d
	})
}
