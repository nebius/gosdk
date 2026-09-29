package gosdk

import (
	"log/slog"
	"sort"

	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/internal/logging"
)

// Credentials are used to authenticate outgoing gRPC requests.
// To disable authentication for a specific request use the [auth.Disable] call option.
type Credentials interface {
	credentials()
}

// NoCredentials is the default [Credentials] that disables authentication.
func NoCredentials() Credentials {
	return credsNoCreds{}
}

// IAMToken is a [Credentials] that uses a predefined token for authentication.
func IAMToken(token string) Credentials {
	return credsTokener{tokener: auth.StaticBearerToken(token)}
}

// CustomTokener uses a custom [auth.BearerTokener] implementation for authentication.
// Structured logs use [slog.LogValuer.LogValue], then [fmt.Stringer.String], or only the implementation type.
// LogValue and String implementations must sanitize their own sensitive values.
func CustomTokener(tokener auth.BearerTokener) Credentials {
	return credsTokener{tokener: tokener}
}

// CustomAuthenticator uses a custom [auth.Authenticator] implementation.
// Structured logs use [slog.LogValuer.LogValue], then [fmt.Stringer.String], or only the implementation type.
// LogValue and String implementations must sanitize their own sensitive values.
func CustomAuthenticator(auth auth.Authenticator) Credentials {
	return credsAuthenticator{auth: auth}
}

// ServiceAccount is the same as [ServiceAccountReader] but with constant [auth.ServiceAccount].
func ServiceAccount(account auth.ServiceAccount) Credentials {
	return credsServiceAccount{reader: auth.StaticServiceAccount(account)}
}

// ServiceAccountReader authorizes gRPC requests using [auth.ServiceAccount].
// It receives an [auth.BearerToken] from the IAM by exchanging a JWT.
// The JWT is signed using the private key of the service account.
//
// The [SDK] ensures a continuously valid bearer token by caching the current token
// and asynchronously requesting a new one before expiration.
//
// The SDK can log the reader.
// Structured logs use [slog.LogValuer.LogValue], then [fmt.Stringer.String], or only the reader type.
// LogValue and String implementations must sanitize their own sensitive values.
//
// Note: the reader is used only once as it is wrapped with [auth.CachedServiceAccount].
func ServiceAccountReader(reader auth.ServiceAccountReader) Credentials {
	switch reader.(type) {
	case *auth.CachedServiceAccount, auth.StaticServiceAccount:
		return credsServiceAccount{reader: reader}
	default:
		cached := auth.NewCachedServiceAccount(reader)
		return credsServiceAccount{reader: cached}
	}
}

// OneOfCredentials allows you to use different credentials for different requests.
// Exactly one [auth.Selector] from creds map must be added to call options to choose one.
//
// You can use predefined [auth.Base] and [auth.Propagate] selectors as well as [auth.Select] with custom name.
func OneOfCredentials(creds map[auth.Selector]Credentials) Credentials {
	return credsOneOf(creds)
}

// PropagateAuthorizationHeader uses [auth.AuthorizationHeader] from the incoming grpc metadata
// and puts it into the outgoing metadata.
func PropagateAuthorizationHeader() Credentials {
	return credsPropagate{}
}

type (
	credsNoCreds        struct{}
	credsTokener        struct{ tokener auth.BearerTokener }
	credsAuthenticator  struct{ auth auth.Authenticator }
	credsServiceAccount struct{ reader auth.ServiceAccountReader }
	credsOneOf          map[auth.Selector]Credentials
	credsPropagate      struct{}
)

func (credsNoCreds) LogValue() slog.Value {
	return slog.GroupValue(slog.String("type", "NoCredentials"))
}
func (c credsTokener) LogValue() slog.Value {
	return slog.GroupValue(slog.String("type", "Tokener"), logging.Object("tokener", c.tokener))
}
func (c credsAuthenticator) LogValue() slog.Value {
	return slog.GroupValue(slog.String("type", "Authenticator"), logging.Object("authenticator", c.auth))
}
func (c credsServiceAccount) LogValue() slog.Value {
	return slog.GroupValue(slog.String("type", "ServiceAccountReader"), logging.Object("reader", c.reader))
}

func (c credsOneOf) LogValue() slog.Value {
	options := make([]slog.Attr, 0, len(c))
	for selector, credentials := range c {
		options = append(options, logging.Object(selector.Name, credentials))
	}
	sort.Slice(options, func(i, j int) bool { return options[i].Key < options[j].Key })
	return slog.GroupValue(
		slog.String("type", "OneOfCredentials"),
		slog.Int("count", len(c)),
		slog.Attr{Key: "options", Value: slog.GroupValue(options...)},
	)
}
func (credsPropagate) LogValue() slog.Value {
	return slog.GroupValue(slog.String("type", "PropagateAuthorizationHeader"))
}

func (credsNoCreds) credentials()        {}
func (credsTokener) credentials()        {}
func (credsAuthenticator) credentials()  {}
func (credsServiceAccount) credentials() {}

func (credsOneOf) credentials()     {}
func (credsPropagate) credentials() {}
