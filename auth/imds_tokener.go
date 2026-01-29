package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// IMDSTokenizer is a [BearerTokener] that serves token from IMDS HTTP endpoint,
// requesting it every time. The endpoint is expected to return JSON-encoded [Token].
type IMDSTokenizer struct {
	endpoint    string
	client      *http.Client
	baseBackoff time.Duration
	maxAttempts int
}

var _ BearerTokener = (*IMDSTokenizer)(nil)

// NewIMDSTokenizer returns a [BearerTokener] that serves token from IMDS HTTP endpoint,
// requesting it every time.
func NewIMDSTokenizer(endpoint string, maxAttempts int, baseBackoff time.Duration) (*IMDSTokenizer, error) {
	if endpoint == "" {
		return nil, errors.New("empty IMDS endpoint")
	}
	return &IMDSTokenizer{
		endpoint:    endpoint,
		client:      http.DefaultClient,
		maxAttempts: maxAttempts,
		baseBackoff: baseBackoff,
	}, nil
}

func (t *IMDSTokenizer) BearerToken(ctx context.Context) (BearerToken, error) {
	for attempt := range t.maxAttempts {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, t.endpoint, nil)
		if err != nil {
			return BearerToken{}, fmt.Errorf("build IMDS token request: %w", err)
		}
		req.Header.Set("Metadata", "true")
		resp, err := t.client.Do(req)
		if err != nil {
			return BearerToken{}, fmt.Errorf("request IMDS token: %w", err)
		}
		statusCode := resp.StatusCode
		if statusCode == http.StatusOK {
			var token Token
			err = json.NewDecoder(resp.Body).Decode(&token)
			_ = resp.Body.Close()
			if err != nil {
				return BearerToken{}, fmt.Errorf("decode IMDS token response: %w", err)
			}
			return BearerToken{
				Token:     token.AccessToken,
				ExpiresAt: token.ExpiresAt,
			}, nil
		}

		if !shouldRetryIMDS(statusCode) || attempt == t.maxAttempts-1 {
			body, _ := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			return BearerToken{}, fmt.Errorf("IMDS token request failed: %s: %s", resp.Status, string(body))
		}
		backoff := t.baseBackoff * time.Duration(attempt+1)
		select {
		case <-ctx.Done():
			return BearerToken{}, ctx.Err()
		case <-time.After(backoff):
		}
	}
	return BearerToken{}, errors.New("failed to get IMDS token")
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
