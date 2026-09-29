package gosdk_test

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/config"
)

func TestTokenLogging(t *testing.T) {
	t.Parallel()
	const raw = "ne1payload.secret-signature"
	token := auth.NewStaticBearerToken(raw)
	const safe = "ne1payload.**"
	expires := time.Date(2030, 1, 2, 3, 4, 5, 0, time.UTC)
	for _, tc := range []struct {
		name   string
		value  any
		fields map[string]any
	}{
		{"direct", token, map[string]any{"type": "StaticBearerToken", "token": safe}},
		{"bearer", auth.BearerToken{Token: raw, ExpiresAt: expires}, map[string]any{"token": safe, "expires_at": "2030-01-02T03:04:05Z"}},
		{"bearer-pointer", &auth.BearerToken{Token: raw}, map[string]any{"token": safe}},
		{"empty", auth.BearerToken{}, map[string]any{"token": ""}},
		{"iam", gosdk.IAMToken(raw), map[string]any{"type": "Tokener", "tokener.type": "StaticBearerToken", "tokener.token": safe}},
		{"custom", gosdk.CustomTokener(token), map[string]any{"tokener.token": safe}},
		{"authenticator", auth.NewAuthenticatorFromBearerTokener(token), map[string]any{"type": "AuthenticatorFromBearerTokener", "tokener.token": safe}},
		{"custom-authenticator", gosdk.CustomAuthenticator(auth.NewAuthenticatorFromBearerTokener(token)), map[string]any{"type": "Authenticator", "authenticator.tokener.token": safe}},
		{"instrumented", auth.NewStaticTokener(raw), map[string]any{"type": "InstrumentedBearerTokener", "tokener.token": safe}},
		{"named", auth.NewNameWrapper("test", token), map[string]any{"type": "NameWrapper", "name": "test", "tokener_type": "static", "tokener.token": safe}},
		{"named-override", auth.NewTypedNameWrapper("test", "override", token), map[string]any{"tokener_type": "override", "tokener.token": safe}},
		{"named-custom", auth.NewNameWrapper("test", nil), map[string]any{"tokener_type": "custom"}},
		{"nested", gosdk.CustomTokener(auth.NewInstrumentedBearerTokener(auth.NewNameWrapper("test", auth.NewStaticTokener(raw)))), map[string]any{"tokener.tokener.name": "test", "tokener.tokener.tokener.tokener.token": safe}},
		{"one-of", gosdk.OneOfCredentials(map[auth.Selector]gosdk.Credentials{auth.Base: gosdk.IAMToken(raw), auth.Propagate: gosdk.PropagateAuthorizationHeader()}), map[string]any{"type": "OneOfCredentials", "count": float64(2), "options.base.tokener.token": safe, "options.propagate.type": "PropagateAuthorizationHeader"}},
		{"empty-one-of", gosdk.OneOfCredentials(nil), map[string]any{"type": "OneOfCredentials", "count": float64(0)}},
		{"none", gosdk.NoCredentials(), map[string]any{"type": "NoCredentials"}},
		{"propagate", gosdk.PropagateAuthorizationHeader(), map[string]any{"type": "PropagateAuthorizationHeader"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			checkStructuredLog(t, tc.value, tc.fields)
		})
	}
}

func checkStructuredLog(t *testing.T, value any, fields map[string]any) {
	t.Helper()
	for _, jsonOutput := range []bool{false, true} {
		var output bytes.Buffer
		var handler slog.Handler = slog.NewTextHandler(&output, nil)
		if jsonOutput {
			handler = slog.NewJSONHandler(&output, nil)
		}
		slog.New(handler).Info("using credentials", "credentials", value)
		for _, secret := range []string{"secret-signature", "secret-private-key", "987654321"} {
			require.NotContains(t, output.String(), secret)
		}
		if !jsonOutput {
			for path := range fields {
				assert.Contains(t, output.String(), "credentials."+path+"=")
			}
			continue
		}
		checkJSONLogFields(t, output.Bytes(), fields)
	}
}

func checkJSONLogFields(t *testing.T, output []byte, fields map[string]any) {
	t.Helper()
	var record map[string]any
	require.NoError(t, json.Unmarshal(output, &record))
	for path, want := range fields {
		got := record["credentials"]
		for key := range strings.SplitSeq(path, ".") {
			object, ok := got.(map[string]any)
			require.True(t, ok, "%s is not a structured object: %s", path, output)
			got = object[key]
		}
		assert.Equal(t, want, got, "field %s", path)
	}
}

func TestProviderLogging(t *testing.T) {
	t.Parallel()
	token := auth.NewStaticBearerToken("ne1payload.secret-signature")
	named := auth.NewNameWrapper("test", token)
	multiprocess, err := auth.NewMultiprocessSyncTokener(token)
	require.NoError(t, err)
	file, err := auth.NewFileTokener("token.txt", time.Minute)
	require.NoError(t, err)
	imds, err := auth.NewIMDSTokenizer("http://imds.test")
	require.NoError(t, err)
	for _, tc := range []struct {
		name   string
		value  any
		fields map[string]any
	}{
		{"cache", auth.NewCachedBearerTokener(token), map[string]any{"type": "CachedBearerTokener", "tokener.token": "ne1payload.**"}},
		{"service-cache", auth.NewCachedTokener(token), map[string]any{"type": "CachedServiceTokener", "tokener.token": "ne1payload.**"}},
		{"sync", auth.NewInAppSyncTokener(token), map[string]any{"type": "InAppSyncTokener", "tokener.token": "ne1payload.**"}},
		{"multiprocess", multiprocess, map[string]any{"type": "MultiprocessSyncTokener", "tokener.token": "ne1payload.**"}},
		{"impersonated", auth.NewExchangeImpersonatedBearerTokener("account", token, nil), map[string]any{"service_account_id": "account", "tokener.token": "ne1payload.**"}},
		{"file-cache", auth.NewFileCacheTokener(named), map[string]any{"name": "test", "tokener.tokener.token": "ne1payload.**"}},
		{"async-file-cache", auth.NewAsynchronouslyRenewableFileCacheTokener(named), map[string]any{"name": "test", "tokener.tokener.token": "ne1payload.**"}},
		{"pure-file-cache", auth.NewPureFileCachedTokener("test"), map[string]any{"name": "test"}},
		{"file", file, map[string]any{"tokener.path": "token.txt"}},
		{"imds", imds, map[string]any{"endpoint": "http://imds.test"}},
		{"federation", auth.NewFederationTokener("client", "https://federation.test", "federation", "profile"), map[string]any{"client_id": "client", "federation_id": "federation", "profile": "profile"}},
		{"exchange", auth.NewExchangeableBearerTokener(auth.NewServiceAccountExchangeTokenRequester(auth.NewPrivateKeyParser([]byte("secret-private-key"), "key", "account")), nil), map[string]any{"requester.account.service_account_id": "account"}},
		{"federated-static", auth.NewFederatedCredentialsTokenRequester("account", auth.NewStaticFederatedCredentialsReader(auth.FederatedCredentials("ne1payload.secret-signature"))), map[string]any{"service_account_id": "account", "reader.credentials.token": "ne1payload.**"}},
		{"federated-file", auth.NewFileFederatedCredentialsReader("credentials.json"), map[string]any{"path": "credentials.json"}},
	} {
		t.Run(tc.name, func(t *testing.T) { t.Parallel(); checkStructuredLog(t, tc.value, tc.fields) })
	}
}

func TestWrapperFormatting(t *testing.T) {
	t.Parallel()
	token := auth.NewStaticBearerToken("ne1payload.secret-signature")
	for _, value := range []any{auth.NewStaticTokener(string(token)), auth.NewNameWrapper("test", token), auth.NewAuthenticatorFromBearerTokener(token)} {
		for _, format := range []string{"%v", "%+v", "%s"} {
			got := fmt.Sprintf(format, value)
			assert.NotContains(t, got, "secret-signature")
			assert.Contains(t, got, "ne1payload.**")
		}
	}
	for _, tc := range []struct {
		wrapper *auth.NameWrapper
		typ     string
	}{
		{auth.NewNameWrapper("test", token), "static"},
		{auth.NewTypedNameWrapper("test", "override", token), "override"},
		{auth.NewNameWrapper("test", nil), "custom"},
	} {
		assert.Contains(t, tc.wrapper.String(), fmt.Sprintf("type=%q", tc.typ))
	}
}

func TestServiceAccountLogging(t *testing.T) {
	t.Parallel()
	account := auth.ServiceAccount{
		ServiceAccountID: "account",
		PublicKeyID:      "key",
		PrivateKey:       &rsa.PrivateKey{D: big.NewInt(987654321)},
	}
	for _, tc := range []struct {
		name   string
		value  any
		prefix string
	}{
		{"account", account, ""},
		{"static", auth.NewStaticServiceAccount(account), ""},
		{"credentials", gosdk.ServiceAccount(account), "reader."},
		{"parser", auth.NewPrivateKeyParser([]byte("secret-private-key"), "key", "account"), ""},
		{"cached", gosdk.ServiceAccountReader(auth.NewPrivateKeyParser([]byte("secret-private-key"), "key", "account")), "reader.reader."},
		{"file", auth.NewPrivateKeyFileParser(nil, "key.pem", "key", "account"), ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			checkStructuredLog(
				t,
				tc.value,
				map[string]any{tc.prefix + "service_account_id": "account", tc.prefix + "public_key_id": "key"},
			)
		})
	}
	checkStructuredLog(
		t,
		auth.NewServiceAccountCredentialsFileParser(nil, "credentials.json"),
		map[string]any{"path": "credentials.json"},
	)
}

func TestProfileLogging(t *testing.T) {
	t.Parallel()
	checkStructuredLog(
		t,
		&config.Profile{Name: "test", Endpoint: "example.com", PrivateKey: "secret-private-key"},
		map[string]any{"name": "test", "endpoint": "example.com"},
	)
}

func TestStaticBearerTokenFormatting(t *testing.T) {
	t.Parallel()
	for _, raw := range []string{"", "ne1payload.secret-signature", "v1.payload.secret-signature", "opaque-token", "v0.short"} {
		t.Run(raw, func(t *testing.T) {
			t.Parallel()
			token := auth.NewStaticBearerToken(raw)
			// Formatting follows the existing BearerToken sanitizer policy.
			safe := (auth.BearerToken{Token: raw}).LogValue().Group()[0].Value.String()
			for _, tc := range []struct {
				value any
				want  string
			}{
				{token, fmt.Sprintf("StaticBearerToken(token=%s)", safe)},
				{auth.BearerToken{Token: raw}, (auth.BearerToken{Token: raw}).String()},
			} {
				for _, format := range []string{"%v", "%+v", "%s", "%#v", "%d", "%q", "%x"} {
					assert.Equal(t, tc.want, fmt.Sprintf(format, tc.value), "format %s", format)
				}
			}
			bearer, err := token.BearerToken(t.Context())
			require.NoError(t, err)
			assert.Equal(t, raw, bearer.Token, "logging must preserve the authentication token")
		})
	}
}

type customLoggingProvider interface {
	auth.BearerTokener
	auth.ServiceAccountReader
	auth.FederatedCredentialsReader
	auth.ExchangeTokenRequester
}

type opaqueLoggingProvider struct {
	customLoggingProvider
	Secret string
}

func (p opaqueLoggingProvider) Format(state fmt.State, _ rune) { _, _ = fmt.Fprint(state, p.Secret) }

func (p opaqueLoggingProvider) MarshalJSON() ([]byte, error) { return json.Marshal(p.Secret) }

type stringLoggingProvider struct {
	opaqueLoggingProvider
}

func (stringLoggingProvider) String() string { return "safe-provider" }

type structuredLoggingProvider struct {
	stringLoggingProvider
}

func (structuredLoggingProvider) LogValue() slog.Value {
	return slog.GroupValue(slog.String("kind", "structured-provider"))
}

func TestCustomCredentialLogging(t *testing.T) {
	t.Parallel()
	opaque := opaqueLoggingProvider{Secret: "secret-private-key"}
	stringer := stringLoggingProvider{opaqueLoggingProvider: opaque}
	for _, tc := range []struct {
		name   string
		value  customLoggingProvider
		suffix string
		want   any
	}{
		{"opaque", opaque, "", "gosdk_test.opaqueLoggingProvider"},
		{"string", stringer, "", "safe-provider"},
		{"structured", structuredLoggingProvider{stringLoggingProvider: stringer}, ".kind", "structured-provider"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for _, wrapped := range []struct {
				name  string
				value any
				path  string
			}{
				{"custom", gosdk.CustomTokener(tc.value), "tokener"},
				{"named", auth.NewNameWrapper("custom", tc.value), "tokener"},
				{"instrumented", auth.NewInstrumentedBearerTokener(tc.value), "tokener"},
				{"authenticator", auth.NewAuthenticatorFromBearerTokener(tc.value), "tokener"},
				{"cached", auth.NewCachedBearerTokener(tc.value), "tokener"},
				{"service-account", gosdk.ServiceAccountReader(tc.value), "reader.reader"},
				{"federated", auth.NewFederatedCredentialsTokenRequester("account", tc.value), "reader"},
				{"exchange", auth.NewExchangeableBearerTokener(tc.value, nil), "requester"},
			} {
				t.Run(wrapped.name, func(t *testing.T) {
					t.Parallel()
					checkStructuredLog(t, wrapped.value, map[string]any{wrapped.path + tc.suffix: tc.want})
				})
			}
		})
	}
}
