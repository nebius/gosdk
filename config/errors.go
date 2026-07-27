package config

import (
	"maps"
	"slices"

	"github.com/nebius/gosdk/internal/configerror"
)

type Error = configerror.Error

func NewError(err error) *Error {
	return configerror.New(err)
}

type MissingConfigError struct {
	err error
}

func NewMissingConfigError(err error) *MissingConfigError {
	if err == nil {
		return nil
	}
	return &MissingConfigError{
		err: err,
	}
}

func (e *MissingConfigError) Error() string {
	return "missing configuration: " + e.err.Error()
}

func (e *MissingConfigError) Unwrap() error {
	return e.err
}

type GetProfileError struct {
	err               error
	availableProfiles []string
}

func NewGetProfileError(err error, profiles ProfilesConfig) *GetProfileError {
	if err == nil {
		return nil
	}
	return &GetProfileError{
		err:               err,
		availableProfiles: slices.Collect(maps.Keys(profiles)),
	}
}

func (e *GetProfileError) Error() string {
	return "get profile: " + e.err.Error()
}

func (e *GetProfileError) Unwrap() error {
	return e.err
}

func (e *GetProfileError) AvailableProfiles() []string {
	return e.availableProfiles
}
