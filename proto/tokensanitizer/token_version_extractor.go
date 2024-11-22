package tokensanitizer

import (
	"strings"
)

// TokenVersionExtractor is an interface for extracting token versions.
type TokenVersionExtractor interface {
	// Extract TokenVersion and a boolean indicating if the version was recognized for given token.
	Extract(token string) (TokenVersion, bool)
}

type DefaultTokenVersionExtractor struct {
	versions map[string]TokenVersion
}

func NewDefaultTokenVersionExtractor(versions map[string]TokenVersion) *DefaultTokenVersionExtractor {
	return &DefaultTokenVersionExtractor{versions: versions}
}

// Extract returns the TokenVersion and true if recognized, or an empty TokenVersion and false if not.
func (e *DefaultTokenVersionExtractor) Extract(token string) (TokenVersion, bool) {
	for _, version := range e.versions {
		if strings.HasPrefix(token, version.Prefix) {
			return version, true
		}
	}
	return TokenVersion{}, false
}
