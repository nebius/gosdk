package tokensanitizer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebius/gosdk/proto/tokensanitizer"
)

func TestAccessTokenSanitizer(t *testing.T) {
	t.Parallel()
	sanitizer := tokensanitizer.AccessTokenSanitizer()

	testCases := []struct {
		name      string
		input     string
		expected  string
		supported bool
	}{
		{"Empty input", "", "", false},
		{"V0 short token", "v0.short", "v0.short", true},
		{"V0 long token", "v0.somelongpayloadhere", "v0.somelongpayload**", true},
		{"NE1 valid token", "ne1somepayload.signaturehere", "ne1somepayload.**", true},
		{"NE1 invalid token", "ne1invalidtokens", "ne1invalidtoken**", false},
		{"V1 valid token", "v1.somepayload.signaturehere", "v1.somepayload.**", true},
		{"V1 invalid token", "v1.invalidtokens", "v1.invalidtoken**", false},
		{"Unrecognized short", "unknown", "unknown**", false},
		{"Unrecognized long", "unrecognizedtoken", "unrecognizedtok**", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			result := sanitizer.Sanitize(testCase.input)
			assert.Equal(t, testCase.supported, sanitizer.IsSupported(testCase.input))
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestCredentialsTokenSanitizer(t *testing.T) {
	t.Parallel()
	sanitizer := tokensanitizer.CredentialsSanitizer()

	testCases := []struct {
		name      string
		input     string
		expected  string
		supported bool
	}{
		{"Empty input", "", "", false},
		{"V0 long token", "v0.somelongpayloadhere", "v0.somelongpayload**", true},
		{"NE1 valid token", "ne1somepayload.signaturehere", "ne1somepayload.**", true},
		{"NE1 invalid token", "ne1invalidtokens", "ne1invalidtoken**", false},
		{"V1 valid token", "v1.somepayload.signaturehere", "v1.somepayload.**", true},
		{"V1 invalid token", "v1.invalidtokens", "v1.invalidtoken**", false},
		{"Unrecognized short", "unknown", "unknown**", false},
		{"JWT valid token", "eyJfast.payload.signature", "eyJfast.payload.**", true},
		{"JWT with extra info", "eyJfast.payload.signature.aftersignature", "eyJfast.payload.**.aftersignature", true},
		{"JWT invalid token", "eyJfast.payload", "eyJfast.payload**", false},
		{"Delegation valid token", "nd1payload.signature", "nd1payload.**", true},
		{"Delegation invalid token", "nd1payload", "nd1payload**", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			result := sanitizer.Sanitize(testCase.input)
			assert.Equal(t, testCase.supported, sanitizer.IsSupported(testCase.input))
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestCustomTokenVersion(t *testing.T) {
	t.Parallel()
	versions := map[string]tokensanitizer.TokenVersion{
		"V0":  {Prefix: "v0.", Delimiter: ".", SignaturePosition: tokensanitizer.NoSignature, TokenPartsCount: 1},
		"NE1": {Prefix: "ne1", Delimiter: ".", SignaturePosition: 1, TokenPartsCount: 2},
		"V2":  {Prefix: "v2:", Delimiter: "-", SignaturePosition: 2, TokenPartsCount: 3},
	}
	sanitizer := tokensanitizer.NewTokenSanitizerWithCustomExtractor(
		tokensanitizer.NewDefaultTokenVersionExtractor(versions),
	)

	input := "v2:payload-moredata-signature"
	expected := "v2:payload-moredata-**"
	result := sanitizer.Sanitize(input)
	assert.True(t, sanitizer.IsSupported(input))
	assert.Equal(t, expected, result)
}
