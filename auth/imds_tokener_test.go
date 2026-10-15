package auth_test

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk/auth"
)

func TestIMDSTokenizerErrorIncludesBoundedResponseBody(t *testing.T) {
	t.Parallel()

	bodies := []struct {
		name      string
		body      string
		truncated bool
	}{
		{name: "empty"},
		{name: "short", body: "endpoint denied"},
		{name: "limit", body: strings.Repeat("x", 64<<10)},
		{name: "large", body: "endpoint denied: " + strings.Repeat("x", 70<<10) + "hidden-tail", truncated: true},
	}
	for _, status := range []int{http.StatusBadRequest, http.StatusInternalServerError} {
		for _, tc := range bodies {
			t.Run(strconv.Itoa(status)+"/"+tc.name, func(t *testing.T) {
				t.Parallel()

				var readers []*strings.Reader
				var closers []*trackingReadCloser
				client := &http.Client{
					Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
						reader := strings.NewReader(tc.body)
						body := &trackingReadCloser{reader: reader}
						readers = append(readers, reader)
						closers = append(closers, body)
						return &http.Response{
							StatusCode: status,
							Status:     strconv.Itoa(status) + " custom-status-text",
							Body:       body,
						}, nil
					}),
				}
				tokener, err := auth.NewIMDSTokenizer(
					"http://metadata.nebius.internal/v1/iam/sa/token",
					auth.WithHTTPClient(client),
					auth.WithHTTPMaxAttempts(3),
					auth.WithHTTPBaseBackoff(0),
				)
				require.NoError(t, err)

				_, err = tokener.BearerToken(context.Background())
				require.Error(t, err)
				message := err.Error()
				prefix := "IMDS token request failed: HTTP " + strconv.Itoa(status) + ": "
				if tc.truncated {
					require.True(t, strings.HasPrefix(message, prefix+"endpoint denied: "))
					require.True(t, strings.HasSuffix(message, "... (truncated)"))
					require.NotContains(t, message, "hidden-tail")
					require.Less(t, len(message), (64<<10)+100)
				} else {
					require.Equal(t, prefix+tc.body, message)
				}
				require.NotContains(t, message, "custom-status-text")
				attempts := 1
				if status == http.StatusInternalServerError {
					attempts = 3
				}
				require.Len(t, readers, attempts)
				for _, closer := range closers {
					require.True(t, closer.closed)
				}
				require.Zero(t, readers[len(readers)-1].Len(), "drain the response for connection reuse")
			})
		}
	}
}

func TestIMDSTokenizerReadErrorPreservesHTTPStatus(t *testing.T) {
	t.Parallel()

	for _, status := range []int{http.StatusBadRequest, http.StatusInternalServerError} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			t.Parallel()

			body := &trackingReadCloser{reader: iotest.ErrReader(io.ErrUnexpectedEOF)}
			client := &http.Client{
				Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: status, Body: body}, nil
				}),
			}
			tokener, err := auth.NewIMDSTokenizer(
				"http://metadata.nebius.internal/v1/iam/sa/token",
				auth.WithHTTPClient(client),
				auth.WithHTTPMaxAttempts(1),
			)
			require.NoError(t, err)

			_, err = tokener.BearerToken(context.Background())
			require.EqualError(
				t, err,
				"IMDS token request failed: HTTP "+strconv.Itoa(status)+"; could not read response body",
			)
			require.True(t, body.closed)
		})
	}
}
