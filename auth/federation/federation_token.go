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
		logger:             logger,
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

	return auth.BearerToken{
		Token:     res.AccessToken,
		ExpiresAt: expiry,
	}, nil
}

func (f *FederationToken) HandleError(context.Context, auth.BearerToken, error) error {
	return nil // token can be renewed
}

func (f *FederationToken) Name() string {
	return fmt.Sprintf("federation/%s/%s/%s", f.federationEndpoint, f.federationID, f.profileName)
}
