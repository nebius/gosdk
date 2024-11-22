package serviceerror

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type Error struct {
	Wrapped error
	Details []Detail
}

func collectServiceErrors(details []*anypb.Any) []Detail {
	res := make([]Detail, 0, len(details))
	for _, detail := range details {
		if d, err := NewDetailFromAny(detail); err == nil {
			res = append(res, d)
		}
	}
	return res
}

func FromError(err error, extraDetails ...*anypb.Any) error {
	if err == nil {
		return nil
	}

	var alreadyError *Error
	if errors.As(err, &alreadyError) {
		return err
	}

	var details []Detail
	if stat, ok := status.FromError(err); ok {
		details = collectServiceErrors(stat.Proto().GetDetails())
	}
	details = append(details, collectServiceErrors(extraDetails)...)

	if len(details) == 0 {
		return err
	}

	return &Error{
		Wrapped: err,
		Details: details,
	}
}

func (e *Error) Error() string {
	return e.Wrapped.Error()
}

func (e *Error) Cause() string {
	errorStrings := make([]string, 0, len(e.Details))
	for _, se := range e.Details {
		errorStrings = append(errorStrings, addIndent(se.String(), "  "))
	}
	if len(errorStrings) > 1 {
		return fmt.Sprintf(
			"caused by service errors:\n%s",
			strings.Join(errorStrings, "\n"),
		)
	}
	return fmt.Sprintf(
		"caused by service error:\n%s",
		errorStrings[0],
	)
}

func (e *Error) Unwrap() error {
	return e.Wrapped
}

func addIndent(s, indent string) string {
	// Split the string into lines
	lines := strings.Split(s, "\n")

	// Prepend the indent to each line
	for i, line := range lines {
		lines[i] = indent + line
	}

	// Join the lines back into a single string
	return strings.Join(lines, "\n")
}
