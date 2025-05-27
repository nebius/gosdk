package serviceerror

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

// Error is a wrapper returned by gRPC services if an original error contains [Detail].
// It also contains request ID if it is present.
type Error struct {
	Wrapped   error
	Details   []Detail
	RequestID string
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

// FromError converts wraps an error with [Error] if necessary.
// The client should not use this function. It is called by built-in interceptor.
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

const (
	requestIDKey = "X-Request-ID"
)

// FromRPCError converts wraps an error with [Error] if necessary.
// It also extracts request ID from metadata if available.
// The client should not use this function. It is called by a built-in interceptor.
func FromRPCError(err error, md *metadata.MD, extraDetails ...*anypb.Any) error {
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

	ret := &Error{
		Wrapped: err,
		Details: details,
	}

	if requestIDs := md.Get(requestIDKey); len(requestIDs) > 0 {
		ret.RequestID = requestIDs[0]
	}

	return ret
}

func (e *Error) Error() string {
	requestInfo := ""
	if e.RequestID != "" {
		requestInfo = fmt.Sprintf("request = %s", e.RequestID)
	}
	cause := e.cause()
	wrapped := e.Wrapped.Error()

	if cause == "" && requestInfo == "" {
		return wrapped
	}

	sep := " "
	if cause != "" || strings.Contains(wrapped, "\n") {
		sep = "\n"
	}

	withRequestInfo := wrapped
	if requestInfo != "" {
		withRequestInfo = wrapped + sep + requestInfo
	}

	if cause != "" {
		return withRequestInfo + "\n" + cause
	}

	return withRequestInfo
}

func (e *Error) cause() string {
	if len(e.Details) == 0 {
		return ""
	}
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
