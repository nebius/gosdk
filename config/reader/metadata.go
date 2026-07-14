package reader

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/nebius/gosdk/config"
	"github.com/nebius/gosdk/config/paths"
	"github.com/nebius/gosdk/constants"
)

const (
	// #nosec G101 -- metadataTokenEndpoint is a URL, not credentials.
	defaultMetadataTokenEndpoint = "http://metadata.nebius.internal/v1/iam/sa/token"
	defaultMetadataProbeTimeout  = 500 * time.Millisecond
	// #nosec G101 -- defaultMetadataTokenFilePath is a file path, not credentials.
	defaultMetadataTokenFilePath = "/mnt/cloud-metadata/token"
	virtualProfileName           = "virtual"
)

type metadataProbe struct {
	tokenEndpoint string
	tokenFilePath string
	httpClient    *http.Client
	timeout       time.Duration
	maxAttempts   int
	baseBackoff   time.Duration
}

type metadataProbeResult int

const (
	metadataProbeUnavailable metadataProbeResult = iota
	metadataProbeAvailable
	metadataProbeRetryable
)

func defaultMetadataProbe() metadataProbe {
	return metadataProbe{
		tokenEndpoint: defaultMetadataTokenEndpoint,
		tokenFilePath: defaultMetadataTokenFilePath,
		httpClient:    http.DefaultClient,
		timeout:       defaultMetadataProbeTimeout,
		maxAttempts:   constants.HTTPMaxAttempts,
		baseBackoff:   constants.HTTPBaseBackoff,
	}
}

func (p metadataProbe) withDefaults() metadataProbe {
	defaults := defaultMetadataProbe()
	if p.tokenEndpoint == "" {
		p.tokenEndpoint = defaults.tokenEndpoint
	}
	if p.tokenFilePath == "" {
		p.tokenFilePath = defaults.tokenFilePath
	}
	if p.httpClient == nil {
		p.httpClient = defaults.httpClient
	}
	if p.timeout == 0 {
		p.timeout = defaults.timeout
	}
	if p.maxAttempts <= 0 {
		p.maxAttempts = defaults.maxAttempts
	}
	if p.baseBackoff == 0 {
		p.baseBackoff = defaults.baseBackoff
	}
	return p
}

func (r *configReader) setVMProfile(ctx context.Context) bool {
	probe := r.metadataProbe.withDefaults()
	profile := &config.Profile{
		Name:     virtualProfileName,
		Endpoint: paths.DefaultAPIEndpoint,
	}
	switch {
	case probe.hasIMDSToken(ctx, r.logger):
		profile.TokenEndpoint = probe.tokenEndpoint
	case probe.hasTokenFile():
		profile.TokenFile = probe.tokenFilePath
	default:
		return false
	}

	r.config.Default = virtualProfileName
	r.config.Profiles[virtualProfileName] = profile
	r.profile = r.config.Profiles[virtualProfileName]
	return true
}

func (p metadataProbe) hasTokenFile() bool {
	if p.tokenFilePath == "" {
		return false
	}
	stat, err := os.Stat(p.tokenFilePath)
	return err == nil && !stat.IsDir()
}

func (p metadataProbe) hasIMDSToken(ctx context.Context, logger *slog.Logger) bool {
	p = p.withDefaults()
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	for attempt := range p.maxAttempts {
		switch p.tryIMDSTokenProbe(ctx, logger, attempt) {
		case metadataProbeAvailable:
			return true
		case metadataProbeUnavailable:
			return false
		case metadataProbeRetryable:
			if attempt == p.maxAttempts-1 {
				return false
			}
			if !p.waitBeforeRetry(ctx, logger, attempt) {
				return false
			}
		}
	}
	return false
}

func (p metadataProbe) tryIMDSTokenProbe(
	ctx context.Context,
	logger *slog.Logger,
	attempt int,
) metadataProbeResult {
	requestCtx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(requestCtx, http.MethodGet, p.tokenEndpoint, nil)
	if err != nil {
		logger.DebugContext(requestCtx, "build metadata token probe request", slog.Any("error", err))
		return metadataProbeUnavailable
	}
	req.Header.Set("Metadata", "true")
	resp, err := p.httpClient.Do(req)
	if err != nil {
		logger.DebugContext(
			requestCtx,
			"metadata token probe request failed",
			slog.Any("error", err),
			slog.Int("attempt", attempt+1),
		)
		return metadataProbeUnavailable
	}
	defer closeMetadataProbeResponseBody(requestCtx, logger, resp)
	if resp.StatusCode != http.StatusOK {
		logger.DebugContext(
			requestCtx,
			"metadata token probe request returned non-OK status",
			slog.Int("status", resp.StatusCode),
			slog.Int("attempt", attempt+1),
		)
		if shouldRetryMetadataProbe(resp.StatusCode) {
			return metadataProbeRetryable
		}
		return metadataProbeUnavailable
	}
	return metadataProbeAvailable
}

func (p metadataProbe) waitBeforeRetry(ctx context.Context, logger *slog.Logger, attempt int) bool {
	backoff := p.baseBackoff * time.Duration(attempt+1)
	if backoff <= 0 {
		return true
	}
	timer := time.NewTimer(backoff)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		logger.DebugContext(ctx, "metadata token probe retry canceled", slog.Any("error", ctx.Err()))
		return false
	case <-timer.C:
		return true
	}
}

func closeMetadataProbeResponseBody(ctx context.Context, logger *slog.Logger, resp *http.Response) {
	if resp.Body == nil {
		return
	}
	if closeErr := resp.Body.Close(); closeErr != nil {
		logger.DebugContext(ctx, "close metadata token probe response body", slog.Any("error", closeErr))
	}
}

func shouldRetryMetadataProbe(statusCode int) bool {
	return statusCode == http.StatusTooManyRequests || statusCode >= http.StatusInternalServerError
}
