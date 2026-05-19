package auth

import (
	"io"
	"time"
)

type federationWriterOption struct {
	writer io.Writer
}

func (o federationWriterOption) apply(t BearerTokener) {
	applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
		ft, ok := t.(*FederationTokener)
		if ok {
			ft.writer = o.writer
		}
		return ok
	})
}

type federationNoBrowserOpenOption struct {
	noBrowserOpen bool
}

func (o federationNoBrowserOpenOption) apply(t BearerTokener) {
	applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
		ft, ok := t.(*FederationTokener)
		if ok {
			ft.noBrowserOpen = o.noBrowserOpen
		}
		return ok
	})
}

type federationAuthTimeoutOption struct {
	timeout time.Duration
}

func (o federationAuthTimeoutOption) apply(t BearerTokener) {
	applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
		ft, ok := t.(*FederationTokener)
		if ok {
			ft.authTimeout = o.timeout
		}
		return ok
	})
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
