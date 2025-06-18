package config

import (
	"fmt"
	"maps"
	"slices"
)

type ConfigError struct {
	err error
}

func NewConfigError(err error) *ConfigError {
	if err == nil {
		return nil
	}
	return &ConfigError{
		err: fmt.Errorf("config: %w", err),
	}
}

func (ce *ConfigError) Error() string {
	return ce.err.Error()
}

func (ce *ConfigError) Unwrap() error {
	return ce.err
}

type MissingConfigError struct {
	err error
}

func NewMissingConfigError(err error) *MissingConfigError {
	return &MissingConfigError{
		err: fmt.Errorf("missing configuration: %w", err),
	}
}

func (m *MissingConfigError) Error() string {
	return m.err.Error()
}

func (m *MissingConfigError) Unwrap() error {
	return m.err
}

type GetProfileError struct {
	err               error
	availableProfiles []string
}

func NewGetProfileError(err error, profiles ProfilesConfig) *GetProfileError {
	return &GetProfileError{
		err:               fmt.Errorf("get profile: %w", err),
		availableProfiles: slices.Collect(maps.Keys(profiles)),
	}
}

func (e *GetProfileError) Error() string {
	return e.err.Error()
}

func (e *GetProfileError) Unwrap() error {
	return e.err
}

func (e *GetProfileError) AvailableProfiles() []string {
	return e.availableProfiles
}
