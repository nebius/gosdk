package mask

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// FieldKey represents a string used as a field key in [Mask] or [FieldPath].
type FieldKey string

var simpleStringPattern = regexp.MustCompile("^[a-zA-Z0-9_]+$")

// isSimpleString checks if the input string matches the simple string pattern.
func isSimpleString(input string) bool {
	return simpleStringPattern.MatchString(input)
}

// Marshal converts a [FieldKey] to its string representation.
// If the [FieldKey] is a simple string, it returns the string directly.
// Otherwise, it marshals the [FieldKey] using [json.Marshal] encoding.
func (f FieldKey) Marshal() (string, error) {
	if isSimpleString(string(f)) {
		return string(f), nil
	}
	res, err := json.Marshal(string(f))
	if err != nil {
		return "", fmt.Errorf("failed to marshal FieldKey: %w", err)
	}
	return string(res), nil
}

// MarshalText marshals [FieldKey] as arbitrary string unlike [FieldKey.Marshal]
// to incorporate it into other marshalers.
func (f FieldKey) MarshalText() ([]byte, error) {
	return json.Marshal(string(f))
}

// UnmarshalText inverts [FieldKey.MarshalText].
func (f *FieldKey) UnmarshalText(text []byte) error {
	var str string
	err := json.Unmarshal(text, &str)
	*f = FieldKey(str)
	return err
}
