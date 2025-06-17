package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/nebius/gosdk/auth"
)

const (
	authTimeout = 10 * time.Minute
)

type FederationToken struct {
	logger             *slog.Logger
	writer             io.Writer
	clientID           string
	federationID       string
	federationEndpoint string
	profileName        string
	noBrowserOpen      bool
}

var _ auth.NamedTokener = (*FederationToken)(nil)

func NewFederationToken(
	logger *slog.Logger,
	clientID string,
	federationEndpoint string,
	federationID string,
	profileName string,
	noBrowserOpen bool,
) *FederationToken {
	return &FederationToken{
		logger:             logger,
		writer:             nil,
		clientID:           clientID,
		federationID:       federationID,
		federationEndpoint: federationEndpoint,
		profileName:        profileName,
		noBrowserOpen:      noBrowserOpen,
	}
}

func NewFederationTokenWithWriter(
	logger *slog.Logger,
	writer io.Writer,
	clientID string,
	federationEndpoint string,
	federationID string,
	profileName string,
	noBrowserOpen bool,
) *FederationToken {
	return &FederationToken{
		logger: logger.With(
			slog.String("name", "federation_token"),
			slog.String("federation_endpoint", federationEndpoint),
			slog.String("federation_id", federationID),
			slog.String("profile_name", profileName),
			slog.String("client_id", clientID),
			slog.Bool("no_browser_open", noBrowserOpen),
		),
		writer:             writer,
		clientID:           clientID,
		federationID:       federationID,
		federationEndpoint: federationEndpoint,
		profileName:        profileName,
		noBrowserOpen:      noBrowserOpen,
	}
}

func (f *FederationToken) BearerToken(
	ctx context.Context,
) (auth.BearerToken, error) {
	ctx, cancel := context.WithTimeout(ctx, authTimeout)
	defer cancel()
	now := time.Now()

	res, err := Authorize(
		ctx,
		f.clientID, f.federationEndpoint, f.federationID,
		f.logger, f.writer, f.noBrowserOpen,
	)
	if err != nil {
		return auth.BearerToken{}, fmt.Errorf("authorize: %w", err)
	}
	expiry := now.Add(time.Duration(res.ExpiresIn) * time.Second)
	ret := auth.BearerToken{
		Token:     res.AccessToken,
		ExpiresAt: expiry,
	}
	f.logger.DebugContext(ctx, "federation token received",
		slog.String("token", ret.String()),
		slog.Any("received_expires_in", res.ExpiresIn),
		slog.Any("start_time", now),
		slog.Any("finish_time", time.Now()),
		slog.Any("acquisition_duration", time.Since(now)),
	)
	return ret, nil
}

func (f *FederationToken) HandleError(context.Context, auth.BearerToken, error) error {
	return nil // token can be renewed
}

func (f *FederationToken) Name() string {
	return fmt.Sprintf("federation/%s/%s/%s", f.federationEndpoint, f.federationID, f.profileName)
}
