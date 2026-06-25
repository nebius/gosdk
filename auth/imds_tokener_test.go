package auth_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/nebius/gosdk/auth"
)

func TestIMDSTokenizerErrorBodyReportsTruncation(t *testing.T) {
	t.Parallel()

	body := strings.Repeat("x", 70<<10) + "tail-marker"
	bodyReader := strings.NewReader(body)
	client := &http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Status:     "400 Bad Request",
				Body:       io.NopCloser(bodyReader),
			}, nil
		}),
	}
	tokener, err := auth.NewIMDSTokenizer(
		"http://metadata.nebius.internal/v1/iam/sa/token",
		auth.WithHTTPClient(client),
		auth.WithHTTPMaxAttempts(1),
	)
	if err != nil {
		t.Fatalf("new imds tokener: %v", err)
	}

	_, err = tokener.BearerToken(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "IMDS response body truncated after") {
		t.Fatalf("error does not report truncation: %v", err)
	}
	if strings.Contains(err.Error(), "tail-marker") {
		t.Fatalf("error contains data beyond truncation limit: %v", err)
	}
	if bodyReader.Len() != 0 {
		t.Fatalf("response body has %d unread bytes, want 0", bodyReader.Len())
	}
}
