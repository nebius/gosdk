package auth

import (
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/nebius/gosdk/proto/tokensanitizer"
)

// BearerToken is a token used in the Bearer authentication scheme.
type BearerToken struct {
	Token     string
	ExpiresAt time.Time // ExpiresAt is the time when the token expires. If it is zero, the token does not expire.
}

var iamTokenSanitizer = tokensanitizer.AccessTokenSanitizer()

type yamlToken struct {
	Token     string `yaml:"token,omitempty"`
	ExpiresAt int64  `yaml:"expires_at,omitempty"` // Unix timestamp
}

func (t BearerToken) String() string {
	if len(t.Token) == 0 {
		return "Token(<empty>)"
	}
	sb := strings.Builder{}
	sb.WriteString("Token(")
	sb.WriteString(iamTokenSanitizer.Sanitize(t.Token))
	if !t.ExpiresAt.IsZero() {
		sb.WriteString(", expires at ")
		sb.WriteString(t.ExpiresAt.String())
		sb.WriteString(", in ")
		sb.WriteString(time.Until(t.ExpiresAt).Round(time.Second).String())
	}
	sb.WriteString(")")
	return sb.String()
}

func (t BearerToken) MarshalYAML() (any, error) {
	ret := yamlToken{
		Token: t.Token,
	}
	ret.ExpiresAt = t.ExpiresAt.Unix()
	if t.ExpiresAt.IsZero() {
		ret.ExpiresAt = 0 // Use 0 if ExpiresAt is not set
	}
	return ret, nil
}

func (t *BearerToken) UnmarshalYAML(node *yaml.Node) error {
	var data yamlToken
	if err := node.Decode(&data); err != nil {
		return err
	}
	t.Token = data.Token
	if data.ExpiresAt != 0 {
		t.ExpiresAt = time.Unix(data.ExpiresAt, 0)
	}
	return nil
}

func (t BearerToken) Equal(other BearerToken) bool {
	return t.Token == other.Token && t.ExpiresAt.Equal(other.ExpiresAt)
}

func (t BearerToken) NotEmpty() bool {
	return t.Token != ""
}

func (t BearerToken) Empty() bool {
	return !t.NotEmpty()
}

func (t BearerToken) IsExpired() bool {
	if t.ExpiresAt.IsZero() {
		return false // If ExpiresAt is not set, we consider the token valid
	}
	return t.ExpiresAt.Before(time.Now())
}
