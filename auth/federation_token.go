package auth

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"time"
)

const federationAuthTimeout = 10 * time.Minute

type federationAuthorizerFunc func(
	ctx context.Context,
	clientID string,
	federationURL string,
	federationID string,
	logger *slog.Logger,
	writer io.Writer,
	noBrowserOpen bool,
) (BearerToken, error)

var federationAuthorizer federationAuthorizerFunc = func( //nolint:gochecknoglobals // temporary bridge to cut the auth <-> federation cycle while deprecated auth/federation handles remain
	context.Context,
	string,
	string,
	string,
	*slog.Logger,
	io.Writer,
	bool,
) (BearerToken, error) {
	return BearerToken{}, errors.New(
		"federation auth implementation is not registered; import nebius.ai/nebo/api/tools/gosdk/auth/federation",
	)
}

// RegisterFederationAuthorizer wires the low-level federation flow implementation
// into the auth package. It is intended for internal package initialization only.
// This bridge exists only while the deprecated federation package constructors
// and aliases still live under auth/federation and we need to cut the
// auth <-> federation dependency cycle.
// After those deprecated entrypoints are removed, drop this registration hook
// and the package-level function variable, remove the federation side-effect
// imports owned by gosdk, and embed this authorization code in the auth-side
// FederationTokener implementation directly.
func RegisterFederationAuthorizer(authorizer federationAuthorizerFunc) {
	if authorizer == nil {
		return
	}
	federationAuthorizer = authorizer
}

type FederationTokener struct {
	logger             *slog.Logger
	writer             io.Writer
	clientID           string
	federationID       string
	federationEndpoint string
	profileName        string
	noBrowserOpen      bool
	authTimeout        time.Duration
}

var _ NamedTokener = (*FederationTokener)(nil)

func NewFederationTokener(
	clientID string,
	federationEndpoint string,
	federationID string,
	profileName string,
	opts ...Option,
) *FederationTokener {
	t := &FederationTokener{
		logger:             slog.New(slog.DiscardHandler),
		writer:             nil,
		clientID:           clientID,
		federationID:       federationID,
		federationEndpoint: federationEndpoint,
		profileName:        profileName,
		noBrowserOpen:      false,
		authTimeout:        federationAuthTimeout,
	}
	applyOptions(t, opts...)
	t.updateLoggerAttributes()
	return t
}

func (f *FederationTokener) SetLogger(logger *slog.Logger) {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	f.logger = logger
	f.updateLoggerAttributes()
}

func (f *FederationTokener) updateLoggerAttributes() {
	f.logger = f.logger.With(
		slog.String("name", "federation_token"),
		slog.String("federation_endpoint", f.federationEndpoint),
		slog.String("federation_id", f.federationID),
		slog.String("profile_name", f.profileName),
		slog.String("client_id", f.clientID),
		slog.Bool("no_browser_open", f.noBrowserOpen),
	)
}

func (f *FederationTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	ctx, cancel := context.WithTimeout(ctx, f.authTimeout)
	defer cancel()
	now := time.Now()

	token, err := federationAuthorizer(
		ctx,
		f.clientID,
		f.federationEndpoint,
		f.federationID,
		f.logger,
		f.writer,
		f.noBrowserOpen,
	)
	if err != nil {
		return BearerToken{}, fmt.Errorf("authorize: %w", err)
	}

	f.logger.DebugContext(ctx, "federation token received",
		slog.String("token", token.String()),
		slog.Any("start_time", now),
		slog.Any("finish_time", time.Now()),
		slog.Any("acquisition_duration", time.Since(now)),
		slog.Any("expires_at", token.ExpiresAt),
	)
	return token, nil
}

func (f *FederationTokener) HandleError(context.Context, BearerToken, error) error {
	return nil
}

func (f *FederationTokener) Name() string {
	return fmt.Sprintf("federation/%s/%s/%s", f.federationEndpoint, f.federationID, f.profileName)
}
