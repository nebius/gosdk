// Package logging provides safe representations of SDK objects.
package logging

import (
	"fmt"
	"log/slog"
)

// Object logs value through LogValue, then String, or only its type.
// Implementations of LogValue and String must sanitize their own sensitive values.
// Evaluation is deferred until the log handler resolves the attribute.
func Object(key string, value any) slog.Attr {
	return slog.Any(key, object{value: value})
}

type object struct {
	value any
}

func (o object) LogValue() slog.Value {
	switch value := o.value.(type) {
	case nil:
		return slog.AnyValue(nil)
	case slog.LogValuer:
		return slog.AnyValue(value)
	case fmt.Stringer:
		return slog.StringValue(value.String())
	default:
		return slog.StringValue(fmt.Sprintf("%T", value))
	}
}
