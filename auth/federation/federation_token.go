package federation

import (
	"io"
	"log/slog"

	"github.com/nebius/gosdk/auth"
)

// Tokener alias to auth.FederationTokener for backwards compatibility.
//
// Deprecated: use [auth.FederationTokener] instead.
type Tokener = auth.FederationTokener

// Deprecated: use [auth.NewFederationTokener] instead.
func NewTokener(
	logger *slog.Logger,
	clientID string,
	federationEndpoint string,
	federationID string,
	profileName string,
	noBrowserOpen bool,
) *Tokener {
	return NewTokenerWithWriter(
		logger,
		nil,
		clientID,
		federationEndpoint,
		federationID,
		profileName,
		noBrowserOpen,
	)
}

// Deprecated: use [auth.NewFederationTokener] instead.
func NewTokenerWithWriter(
	logger *slog.Logger,
	writer io.Writer,
	clientID string,
	federationEndpoint string,
	federationID string,
	profileName string,
	noBrowserOpen bool,
) *Tokener {
	logger.Warn("you're using a deprecated tokener, consider switching to" +
		" auth.NewFederationTokener with options for better retry behavior and easier maintenance." +
		" See the documentation for details.")
	return auth.NewFederationTokener(
		clientID,
		federationEndpoint,
		federationID,
		profileName,
		auth.WithFederationWriter(writer),
		auth.WithLogger(logger),
		auth.WithFederationNoBrowserOpenFlag(noBrowserOpen),
	)
}
