package auth

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/nebius/gosdk/auth/federation"
)

const federationAuthTimeout = 10 * time.Minute

type FederationTokener struct {
	metrics            atomicMetrics
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
var _ MetricsSetter = (*FederationTokener)(nil)
var _ TypedTokener = (*FederationTokener)(nil)

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

func (f *FederationTokener) SetMetrics(metrics Metrics) {
	f.metrics.Store(metrics)
}

func (f *FederationTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	ctx, cancel := context.WithTimeout(ctx, f.authTimeout)
	defer cancel()
	now := time.Now()

	res, err := federation.Authorize(
		ctx,
		f.clientID,
		f.federationEndpoint,
		f.federationID,
		f.logger,
		f.writer,
		f.noBrowserOpen,
	)
	if err != nil {
		f.metrics.tokenAcquireError(ctx, f, time.Since(now), 0)
		return BearerToken{}, fmt.Errorf("authorize: %w", err)
	}
	token := BearerToken{
		Token:     res.AccessToken,
		ExpiresAt: now.Add(time.Duration(res.ExpiresIn) * time.Second),
	}
	f.metrics.tokenAcquireSuccess(ctx, f, time.Since(now), 0, token, time.Now())

	f.logger.DebugContext(ctx, "federation token received",
		slog.String("token", token.String()), // token is sanitized in String()
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

func (f *FederationTokener) Type() string {
	return "federation"
}
