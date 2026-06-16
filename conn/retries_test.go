package conn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nebius/gosdk/conn"
)

func TestIsRetriableNebiusError_UnexpectedHTTP5xxStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "cloudflare 520 grpc unknown status",
			err:  status.Error(codes.Unknown, "unexpected HTTP status code received from server: 520 ()"),
			want: true,
		},
		{
			name: "cloudflare 529 grpc unknown status",
			err: status.Error(
				codes.Unknown,
				"unexpected HTTP status code received from server: 529 (); "+
					`transport: received unexpected content-type "text/html"`,
			),
			want: true,
		},
		{
			name: "generic 52x grpc unknown status",
			err: status.Error(
				codes.Unknown,
				"unexpected HTTP status code received from server: 520 (Bad Gateway)",
			),
			want: true,
		},
		{
			name: "generic 50x grpc unknown status",
			err: status.Error(
				codes.Unknown,
				"unexpected HTTP status code received from server: 500 (Internal Server Error)",
			),
			want: false,
		},
		{
			name: "non 52x grpc unknown status",
			err:  status.Error(codes.Unknown, "unexpected HTTP status code received from server: 499 ()"),
			want: false,
		},
		{
			name: "not an http status",
			err:  status.Error(codes.Unknown, "operation 520 failed"),
			want: false,
		},
		{
			name: "non 52x status with 520 later",
			err: status.Error(
				codes.Unknown,
				"unexpected HTTP status code received from server: 404 (Not Found); body size 520",
			),
			want: false,
		},
		{
			name: "four digit status-like number",
			err:  status.Error(codes.Unknown, "unexpected HTTP status code received from server: 1520"),
			want: false,
		},
		{
			name: "generic status text with later 520",
			err:  status.Error(codes.Unknown, "status: error 40 from proxy 520"),
			want: false,
		},
		{
			name: "expected phrase with non unknown code",
			err:  status.Error(codes.Internal, "unexpected HTTP status code received from server: 520 ()"),
			want: false,
		},
		{
			name: "known retryable grpc status",
			err:  status.Error(codes.Unavailable, "server unavailable"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, conn.IsRetriableNebiusError(tt.err))
		})
	}
}
