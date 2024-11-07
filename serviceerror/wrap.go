package serviceerror

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type WrapError struct {
	wrapped       error
	ServiceErrors []ServiceError
}

func collectServiceErrors(details []*anypb.Any) []ServiceError {
	serviceErrors := make([]ServiceError, 0, len(details))
	for _, detail := range details {
		if se, err := ServiceErrorFromAny(detail); err == nil {
			serviceErrors = append(serviceErrors, se)
		}
	}
	return serviceErrors
}

func FromKnownErrors(err error) error {
	return FromErrorAndAny(err, nil)
}

func FromErrorAndAny(err error, details []*anypb.Any) error {
	if err == nil {
		return nil
	}
	wrapError := &WrapError{}
	if errors.As(err, &wrapError) {
		return err
	}
	var serviceErrors []ServiceError
	if stat, ok := status.FromError(err); ok {
		serviceErrors = collectServiceErrors(stat.Proto().Details)
	}
	serviceErrors = append(serviceErrors, collectServiceErrors(details)...)

	if len(serviceErrors) > 0 {
		return &WrapError{
			wrapped:       err,
			ServiceErrors: serviceErrors,
		}
	}
	return err
}

func (e *WrapError) Error() string {
	return e.wrapped.Error()
}

func (e *WrapError) Cause() string {
	errorStrings := make([]string, 0, len(e.ServiceErrors))
	for _, se := range e.ServiceErrors {
		errorStrings = append(errorStrings, AddIndent(se.Error(), "  "))
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

func (e *WrapError) Unwrap() error {
	return e.wrapped
}

func AddIndent(s, indent string) string {
	// Split the string into lines
	lines := strings.Split(s, "\n")

	// Prepend the indent to each line
	for i, line := range lines {
		lines[i] = indent + line
	}

	// Join the lines back into a single string
	return strings.Join(lines, "\n")
}
