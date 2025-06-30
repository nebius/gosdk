package config

import (
	"golang.org/x/exp/maps"
)

type Error struct {
	err error
}

func NewError(err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		err: err,
	}
}

func (e *Error) Error() string {
	return "config: " + e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
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
		availableProfiles: maps.Keys(profiles),
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
