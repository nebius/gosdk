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
		token, retry, err := t.tryBearerToken(ctx, start, attempt)
		if err != nil {
			if !retry {
				return BearerToken{}, err
			}
		} else {
			return token, nil
		}
		if waitErr := t.waitBeforeRetry(ctx, start, attempt); waitErr != nil {
			return BearerToken{}, waitErr
		}
	}
	t.metrics.tokenAcquireError(ctx, t, time.Since(start), t.maxAttempts)
	return BearerToken{}, errors.New("failed to get IMDS token")
}

func (t *IMDSTokenizer) tryBearerToken(
	ctx context.Context,
	start time.Time,
	attempt int,
) (BearerToken, bool, error) {
	req, err := t.newIMDSRequest(ctx)
	if err != nil {
		t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
		return BearerToken{}, false, fmt.Errorf("build IMDS token request: %w", err)
	}
	resp, err := t.client.Do(req) //nolint:bodyclose // handleIMDSResponse closes every response body.
	if err != nil {
		t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
		return BearerToken{}, false, fmt.Errorf("request IMDS token: %w", err)
	}
	return t.handleIMDSResponse(ctx, resp, start, attempt)
}

func (t *IMDSTokenizer) newIMDSRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, t.endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Metadata", "true")
	return req, nil
}

func (t *IMDSTokenizer) handleIMDSResponse(
	ctx context.Context,
	resp *http.Response,
	start time.Time,
	attempt int,
) (BearerToken, bool, error) {
	if resp.StatusCode == http.StatusOK {
		token, err := t.decodeIMDSTokenResponse(ctx, resp.Body, start, attempt)
		return token, false, err
	}
	if !shouldRetryIMDS(resp.StatusCode) || attempt == t.maxAttempts-1 {
		return BearerToken{}, false, t.imdsErrorResponse(ctx, resp, start, attempt)
	}
	t.closeRetryableIMDSResponseBody(ctx, resp, attempt)
	return BearerToken{}, true, errors.New("retryable IMDS response")
}

func (t *IMDSTokenizer) decodeIMDSTokenResponse(
	ctx context.Context,
	body io.ReadCloser,
	start time.Time,
	attempt int,
) (BearerToken, error) {
	var token Token
	err := json.NewDecoder(body).Decode(&token)
	closeErr := body.Close()
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

func (t *IMDSTokenizer) imdsErrorResponse(
	ctx context.Context,
	resp *http.Response,
	start time.Time,
	attempt int,
) error {
	body, readErr := readLimitedIMDSResponseBody(resp.Body)
	closeErr := resp.Body.Close()
	t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
	if readErr != nil {
		if closeErr != nil {
			readErr = errors.Join(readErr, fmt.Errorf("close IMDS error response body: %w", closeErr))
		}
		return fmt.Errorf("read IMDS error response body: %w", readErr)
	}
	if closeErr != nil {
		t.logger.WarnContext(
			ctx,
			"failed to close IMDS error response body",
			slog.Any("error", closeErr),
			slog.Int("attempt", attempt+1),
			slog.Int("status_code", resp.StatusCode),
		)
	}
	return fmt.Errorf("IMDS token request failed: %s: %s", resp.Status, string(body))
}

func (t *IMDSTokenizer) closeRetryableIMDSResponseBody(ctx context.Context, resp *http.Response, attempt int) {
	closeErr := resp.Body.Close()
	if closeErr != nil {
		t.logger.WarnContext(
			ctx,
			"failed to close retryable IMDS response body",
			slog.Any("error", closeErr),
			slog.Int("attempt", attempt+1),
			slog.Int("status_code", resp.StatusCode),
		)
	}
}

func (t *IMDSTokenizer) waitBeforeRetry(ctx context.Context, start time.Time, attempt int) error {
	backoff := t.baseBackoff * time.Duration(attempt+1)
	select {
	case <-ctx.Done():
		t.metrics.tokenAcquireError(ctx, t, time.Since(start), attempt+1)
		return ctx.Err()
	case <-time.After(backoff):
		return nil
	}
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
	if _, copyErr := io.Copy(io.Discard, body); copyErr != nil {
		return nil, copyErr
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
