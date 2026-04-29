package federation

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/nebius/gosdk/auth"
)

// This registration bridge exists only while the deprecated federation package
// entrypoints remain in this package. Once those deprecated handles are removed,
// the low-level flow should move next to auth.FederationTokener and this file,
// the init hook, and the federation side-effect imports owned by gosdk can all
// be deleted.
// See the comment on RegisterFederationAuthorizer for more details.
//
//nolint:gochecknoinits // temporary bridge registration while deprecated auth/federation handles remain
func init() {
	auth.RegisterFederationAuthorizer(authorizeBearerToken)
}

func authorizeBearerToken(
	ctx context.Context,
	clientID string,
	federationURL string,
	federationID string,
	logger *slog.Logger,
	writer io.Writer,
	noBrowserOpen bool,
) (auth.BearerToken, error) {
	res, err := Authorize(
		ctx,
		clientID,
		federationURL,
		federationID,
		logger,
		writer,
		noBrowserOpen,
	)
	if err != nil {
		return auth.BearerToken{}, err
	}

	now := time.Now()
	return auth.BearerToken{
		Token:     res.AccessToken,
		ExpiresAt: now.Add(time.Duration(res.ExpiresIn) * time.Second),
	}, nil
}
