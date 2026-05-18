package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/nebius/gosdk/constants"
)

const maxIMDSResponseBodyBytes = 64 << 10

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// IMDSTokenizer is a [BearerTokener] that serves token from IMDS HTTP endpoint,
// requesting it every time. The endpoint is expected to return JSON-encoded [Token].
type IMDSTokenizer struct {
	metrics     atomicMetrics
	logger      *slog.Logger
	endpoint    string
	client      *http.Client
	baseBackoff time.Duration
	maxAttempts int
}

var _ BearerTokener = (*IMDSTokenizer)(nil)
var _ LoggerSetter = (*IMDSTokenizer)(nil)
var _ MetricsSetter = (*IMDSTokenizer)(nil)
var _ TypedTokener = (*IMDSTokenizer)(nil)

func WithHTTPMaxAttempts(attempts int) Option {
	return optionFunc(func(t BearerTokener) {
		applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
			imdsTokener, ok := t.(*IMDSTokenizer)
			if ok {
				imdsTokener.maxAttempts = attempts
			}
			return ok
		})
	})
}

func WithHTTPBaseBackoff(backoff time.Duration) Option {
	return optionFunc(func(t BearerTokener) {
		applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
			imdsTokener, ok := t.(*IMDSTokenizer)
			if ok {
				imdsTokener.baseBackoff = backoff
			}
			return ok
		})
	})
}

func WithHTTPClient(client *http.Client) Option {
	return optionFunc(func(t BearerTokener) {
		applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
			imdsTokener, ok := t.(*IMDSTokenizer)
			if ok {
				imdsTokener.client = client
			}
			return ok
		})
	})
}

// NewIMDSTokenizer returns a [BearerTokener] that serves token from IMDS HTTP endpoint,
// requesting it every time.
func NewIMDSTokenizer(endpoint string, opts ...Option) (*IMDSTokenizer, error) {
	if endpoint == "" {
		return nil, errors.New("empty IMDS endpoint")
	}
	t := &IMDSTokenizer{
		endpoint:    endpoint,
		client:      http.DefaultClient,
		maxAttempts: constants.HTTPMaxAttempts,
		baseBackoff: constants.HTTPBaseBackoff,
		logger:      slog.New(slog.DiscardHandler),
	}
	for _, opt := range opts {
		opt.apply(t)
	}
	return t, nil
}

func (t *IMDSTokenizer) SetLogger(logger *slog.Logger) {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	t.logger = logger
}

func (t *IMDSTokenizer) SetMetrics(metrics Metrics) {
	t.metrics.Store(metrics)
}

func (t *IMDSTokenizer) Type() string {
	return "imds"
}

func (t *IMDSTokenizer) BearerToken(ctx context.Context) (BearerToken, error) {
	start := time.Now()
	for attempt := range t.maxAttempts {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, t.endpoint, nil)
		if err != nil {
			t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
			return BearerToken{}, fmt.Errorf("build IMDS token request: %w", err)
		}
		req.Header.Set("Metadata", "true")
		resp, err := t.client.Do(req)
		if err != nil {
			t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
			return BearerToken{}, fmt.Errorf("request IMDS token: %w", err)
		}
		statusCode := resp.StatusCode
		if statusCode == http.StatusOK {
			var token Token
			err = json.NewDecoder(resp.Body).Decode(&token)
			closeErr := resp.Body.Close()
			if err != nil {
				t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
				if closeErr != nil {
					err = errors.Join(err, fmt.Errorf("close IMDS token response body: %w", closeErr))
				}
				return BearerToken{}, fmt.Errorf("decode IMDS token response: %w", err)
			}
			if closeErr != nil {
				t.logger.WarnContext(
					ctx,
					"failed to close IMDS token response body",
					slog.Any("error", closeErr),
					slog.Int("attempt", attempt+1),
				)
			}
			bearerToken := BearerToken{
				Token:     token.AccessToken,
				ExpiresAt: token.ExpiresAt,
			}
			t.metrics.tokenAcquireSuccess(ctx, t, time.Since(start), attempt+1, bearerToken, time.Now())
			return bearerToken, nil
		}

		if !shouldRetryIMDS(statusCode) || attempt == t.maxAttempts-1 {
			body, readErr := readLimitedIMDSResponseBody(resp.Body)
			closeErr := resp.Body.Close()
			t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
			if readErr != nil {
				if closeErr != nil {
					readErr = errors.Join(readErr, fmt.Errorf("close IMDS error response body: %w", closeErr))
				}
				return BearerToken{}, fmt.Errorf("read IMDS error response body: %w", readErr)
			}
			if closeErr != nil {
				t.logger.WarnContext(
					ctx,
					"failed to close IMDS error response body",
					slog.Any("error", closeErr),
					slog.Int("attempt", attempt+1),
					slog.Int("status_code", statusCode),
				)
			}
			return BearerToken{}, fmt.Errorf("IMDS token request failed: %s: %s", resp.Status, string(body))
		}
		closeErr := resp.Body.Close()
		if closeErr != nil {
			t.logger.WarnContext(
				ctx,
				"failed to close retryable IMDS response body",
				slog.Any("error", closeErr),
				slog.Int("attempt", attempt+1),
				slog.Int("status_code", statusCode),
			)
		}
		backoff := t.baseBackoff * time.Duration(attempt+1)
		select {
		case <-ctx.Done():
			t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
			return BearerToken{}, ctx.Err()
		case <-time.After(backoff):
		}
	}
	t.metrics.tokenAcquireError(ctx, t, time.Since(start), t.maxAttempts)
	return BearerToken{}, errors.New("failed to get IMDS token")
}

func readLimitedIMDSResponseBody(body io.Reader) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(body, maxIMDSResponseBodyBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) <= maxIMDSResponseBodyBytes {
		return data, nil
	}
	data = data[:maxIMDSResponseBodyBytes]
	// Keep the diagnostic body bounded, but drain the rest so net/http can reuse the connection.
	if _, err := io.Copy(io.Discard, body); err != nil {
		return nil, err
	}
	return append(data, fmt.Appendf(nil,
		"\n... IMDS response body truncated after %d bytes",
		maxIMDSResponseBodyBytes,
	)...), nil
}

func shouldRetryIMDS(statusCode int) bool {
	return statusCode == http.StatusTooManyRequests || statusCode >= http.StatusInternalServerError
}

func (t *IMDSTokenizer) HandleError(ctx context.Context, token BearerToken, err error) error {
	rt, tokErr := t.BearerToken(ctx)
	if tokErr != nil {
		return errors.Join(tokErr, err)
	}
	if rt.Token != token.Token {
		return nil
	}

	return err // Same token, but it is invalid, fail
}
