package auth

import (
	"io"
	"time"
)

type federationWriterOption struct {
	writer io.Writer
}

func (o federationWriterOption) apply(t BearerTokener) {
	if ft, ok := t.(*FederationTokener); ok {
		ft.writer = o.writer
	}
}

type federationNoBrowserOpenOption struct {
	noBrowserOpen bool
}

func (o federationNoBrowserOpenOption) apply(t BearerTokener) {
	if ft, ok := t.(*FederationTokener); ok {
		ft.noBrowserOpen = o.noBrowserOpen
	}
}

type federationAuthTimeoutOption struct {
	timeout time.Duration
}

func (o federationAuthTimeoutOption) apply(t BearerTokener) {
	if ft, ok := t.(*FederationTokener); ok {
		ft.authTimeout = o.timeout
	}
}

func WithFederationWriter(writer io.Writer) Option {
	return federationWriterOption{writer: writer}
}

func WithFederationNoBrowserOpenFlag(noBrowserOpen bool) Option {
	return federationNoBrowserOpenOption{noBrowserOpen: noBrowserOpen}
}

func WithFederationNoBrowserOpen() Option {
	return WithFederationNoBrowserOpenFlag(true)
}

func WithFederationBrowserOpen() Option {
	return WithFederationNoBrowserOpenFlag(false)
}

func WithFederationAuthTimeout(timeout time.Duration) Option {
	return federationAuthTimeoutOption{timeout: timeout}
}
