package tokensanitizer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebius/gosdk/proto/tokensanitizer"
)

func TestDefaultTokenVersionExtractor(t *testing.T) {
	t.Parallel()
	versions := map[string]tokensanitizer.TokenVersion{
		"V0":  {Prefix: "v0.", Delimiter: ".", SignaturePosition: tokensanitizer.NoSignature, TokenPartsCount: 1},
		"NE1": {Prefix: "ne1", Delimiter: ".", SignaturePosition: 1, TokenPartsCount: 2},
		"V2":  {Prefix: "v2:", Delimiter: "-", SignaturePosition: 2, TokenPartsCount: 2},
	}

	extractor := tokensanitizer.NewDefaultTokenVersionExtractor(versions)

	testCases := []struct {
		name            string
		input           string
		expectedVersion tokensanitizer.TokenVersion
		expectedFound   bool
	}{
		{
			name:            "V0 token",
			input:           "v0.sometokendata",
			expectedVersion: versions["V0"],
			expectedFound:   true,
		},
		{
			name:            "NE1 token",
			input:           "ne1sometokendata.signature",
			expectedVersion: versions["NE1"],
			expectedFound:   true,
		},
		{
			name:            "V2 token",
			input:           "v2:part1-part2-signature",
			expectedVersion: versions["V2"],
			expectedFound:   true,
		},
		{
			name:            "Unrecognized token",
			input:           "unknowntoken",
			expectedVersion: tokensanitizer.TokenVersion{},
			expectedFound:   false,
		},
		{
			name:            "Empty token",
			input:           "",
			expectedVersion: tokensanitizer.TokenVersion{},
			expectedFound:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			version, found := extractor.Extract(testCase.input)
			assert.Equal(t, testCase.expectedFound, found)
			assert.Equal(t, testCase.expectedVersion, version)
		})
	}
}
