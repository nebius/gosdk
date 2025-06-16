package tokensanitizer

import (
	"strings"
)

// MaskString is the string to show instead of real token part.
const MaskString = "**"

// MaxVisiblePayloadLength is the maximum number of characters to show for unrecognized or not signed token.
const MaxVisiblePayloadLength = 15

// NoSignature is used to indicate that a token version has no signature.
const NoSignature = -1

// TokenVersion represents the structure of a token version.
type TokenVersion struct {
	Prefix            string
	Delimiter         string
	SignaturePosition int // -1 if no signature
	TokenPartsCount   int
}

// AccessTokenVersions contains the predefined access token versions.
//
//nolint:gochecknoglobals,mnd // const config
var AccessTokenVersions = map[string]TokenVersion{
	"V0":  {Prefix: "v0.", Delimiter: ".", SignaturePosition: NoSignature, TokenPartsCount: 1},
	"NE1": {Prefix: "ne1", Delimiter: ".", SignaturePosition: 1, TokenPartsCount: 2},
	"V1":  {Prefix: "v1.", Delimiter: ".", SignaturePosition: 2, TokenPartsCount: 3},

	// When changing this, also update PySDK:
	// https://github.com/nebius/pysdk/blob/main/src/nebius/base/token_sanitizer.py
}

// CredentialsVersions contains all possible token types, which is used in app.
//
//nolint:gochecknoglobals,mnd // const config
var CredentialsVersions = map[string]TokenVersion{
	// Copy entries from AccessTokenVersions
	"V0":  AccessTokenVersions["V0"],
	"NE1": AccessTokenVersions["NE1"],
	"V1":  AccessTokenVersions["V1"],

	// Other cred types delegation and jwt
	"DE1": {Prefix: "nd1", Delimiter: ".", SignaturePosition: 1, TokenPartsCount: 2},
	"JWT": {Prefix: "eyJ", Delimiter: ".", SignaturePosition: 2, TokenPartsCount: 3},

	// When changing this, also update PySDK:
	// https://github.com/nebius/pysdk/blob/main/src/nebius/base/token_sanitizer.py
}

type TokenSanitizer struct {
	extractor TokenVersionExtractor
}

func AccessTokenSanitizer() *TokenSanitizer {
	return &TokenSanitizer{
		extractor: NewDefaultTokenVersionExtractor(AccessTokenVersions),
	}
}

func CredentialsSanitizer() *TokenSanitizer {
	return &TokenSanitizer{
		extractor: NewDefaultTokenVersionExtractor(CredentialsVersions),
	}
}

func NewTokenSanitizerWithCustomExtractor(versionExtractor TokenVersionExtractor) *TokenSanitizer {
	return &TokenSanitizer{
		extractor: versionExtractor,
	}
}

// Sanitize sanitizes the input token based on its version, returns sanitized token.
func (ts *TokenSanitizer) Sanitize(token string) string {
	if token == "" {
		return ""
	}

	version, recognized := ts.extractor.Extract(token)
	if !recognized {
		return sanitizeUnrecognized(token)
	}

	tokenParts := strings.Split(token, version.Delimiter)

	if version.SignaturePosition == NoSignature {
		return sanitizeNoSignature(token, version.Prefix)
	}

	if len(tokenParts) <= version.SignaturePosition {
		return sanitizeUnrecognized(token)
	}

	tokenParts[version.SignaturePosition] = MaskString
	return strings.Join(tokenParts, version.Delimiter)
}

// IsSupported check version is supported and token parts count is suitable for this token type.
func (ts *TokenSanitizer) IsSupported(token string) bool {
	version, recognized := ts.extractor.Extract(token)
	if !recognized {
		return false
	}
	tokenParts := strings.Split(token, version.Delimiter)
	return len(tokenParts) >= version.TokenPartsCount
}

// sanitizeNoSignature handles sanitization for tokens without a signature.
func sanitizeNoSignature(token, prefix string) string {
	payload := strings.TrimPrefix(token, prefix)
	if len(payload) <= MaxVisiblePayloadLength {
		return token
	}
	return token[:MaxVisiblePayloadLength+len(prefix)] + MaskString
}

// sanitizeUnrecognized handles sanitization for unrecognized token structures.
func sanitizeUnrecognized(token string) string {
	if len(token) <= MaxVisiblePayloadLength {
		return token + MaskString
	}
	return token[:MaxVisiblePayloadLength] + MaskString
}
