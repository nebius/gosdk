package auth_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk/auth"
)

// deadlineRecorder records the deadline that the cache gives to its tokener.
// It returns a token at once and never calls the wrapped tokener.
type deadlineRecorder struct {
	deadline time.Time
	hasIt    bool
}

func (d *deadlineRecorder) BearerToken(ctx context.Context) (auth.BearerToken, error) {
	d.deadline, d.hasIt = ctx.Deadline()
	return auth.BearerToken{Token: "token"}, nil
}

func (*deadlineRecorder) HandleError(_ context.Context, _ auth.BearerToken, err error) error {
	return err
}

// budgetProvider is a custom tokener that reports its own budget.
type budgetProvider struct {
	deadlineRecorder
	budget time.Duration
	ok     bool
}

func (p *budgetProvider) AcquisitionBudget() (time.Duration, bool) {
	return p.budget, p.ok
}

// forwardingWrapper is a transparent wrapper that forwards the budget and
// options to the wrapped tokener, and records the deadline instead of calling it.
type forwardingWrapper struct {
	deadlineRecorder
	inner auth.BearerTokener
}

func (w *forwardingWrapper) Unwrap() auth.BearerTokener {
	return w.inner
}

func (w *forwardingWrapper) AcquisitionBudget() (time.Duration, bool) {
	if provider, ok := w.inner.(auth.AcquisitionBudgetProvider); ok {
		return provider.AcquisitionBudget()
	}
	return 0, false
}

// opaqueWrapper unwraps for options but does not report a budget on purpose.
type opaqueWrapper struct {
	deadlineRecorder
	inner auth.BearerTokener
}

func (w *opaqueWrapper) Unwrap() auth.BearerTokener {
	return w.inner
}

func newFederation(opts ...auth.Option) *auth.FederationTokener {
	return auth.NewFederationTokener("client", "endpoint", "federation", "profile", opts...)
}

func newImpersonation(actor auth.BearerTokener) auth.BearerTokener {
	return auth.NewExchangeImpersonatedBearerTokener("serviceaccount-target", actor, nil)
}

type recordingTokener interface {
	auth.BearerTokener
	recorded() (time.Time, bool)
}

func (d *deadlineRecorder) recorded() (time.Time, bool) {
	return d.deadline, d.hasIt
}

func TestCachedServiceTokenerAcquireBudget(t *testing.T) {
	t.Parallel()
	const cacheDefault = 5 * time.Second
	const exchangeAllowance = 5 * time.Second
	for _, test := range []struct {
		name    string
		tokener recordingTokener
		opts    []auth.Option
		want    time.Duration
	}{
		{
			name:    "custom provider budget is used",
			tokener: &budgetProvider{budget: 42 * time.Second, ok: true},
			want:    42 * time.Second,
		},
		{
			name:    "explicit cap wins over the provider budget",
			tokener: &budgetProvider{budget: 42 * time.Second, ok: true},
			opts:    []auth.Option{auth.WithCachedTokenerAcquireTimeout(3 * time.Second)},
			want:    3 * time.Second,
		},
		{
			name:    "tokener without the contract gets the default",
			tokener: &deadlineRecorder{},
			want:    cacheDefault,
		},
		{
			name:    "provider without a budget gets the default",
			tokener: &budgetProvider{ok: false},
			want:    cacheDefault,
		},
		{
			name:    "non-positive provider budget gets the default",
			tokener: &budgetProvider{budget: 0, ok: true},
			want:    cacheDefault,
		},
		{
			name:    "wrapper without the contract is not introspected",
			tokener: &opaqueWrapper{inner: newFederation(auth.WithFederationAuthTimeout(7 * time.Second))},
			want:    cacheDefault,
		},
		{
			name:    "federation owns the login timeout",
			tokener: &forwardingWrapper{inner: auth.NewInAppSyncTokener(newFederation(auth.WithFederationAuthTimeout(7 * time.Second)))},
			want:    7 * time.Second,
		},
		{
			name:    "impersonation adds the exchange allowance",
			tokener: &forwardingWrapper{inner: newImpersonation(auth.NewNameWrapper("actor", newFederation(auth.WithFederationAuthTimeout(7*time.Second))))},
			want:    7*time.Second + exchangeAllowance,
		},
		{
			name:    "late federation option through the cache is seen",
			tokener: &forwardingWrapper{inner: auth.NewInAppSyncTokener(newFederation())},
			opts:    []auth.Option{auth.WithFederationAuthTimeout(9 * time.Second)},
			want:    9 * time.Second,
		},
		{
			name:    "nested cache reports its explicit cap",
			tokener: &forwardingWrapper{inner: auth.NewCachedTokener(&deadlineRecorder{}, auth.WithCachedTokenerAcquireTimeout(3*time.Second))},
			want:    3 * time.Second,
		},
		{
			name:    "impersonation over a zero federation timeout keeps the default",
			tokener: &forwardingWrapper{inner: newImpersonation(newFederation(auth.WithFederationAuthTimeout(0)))},
			want:    cacheDefault,
		},
		{
			name:    "impersonation saturates on overflow",
			tokener: &forwardingWrapper{inner: newImpersonation(newFederation(auth.WithFederationAuthTimeout(math.MaxInt64)))},
			want:    math.MaxInt64,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			cached := auth.NewCachedTokener(test.tokener, test.opts...)
			start := time.Now()

			_, err := cached.BearerToken(t.Context())

			require.NoError(t, err)
			deadline, hasIt := test.tokener.recorded()
			require.True(t, hasIt, "acquisition context has no deadline")
			require.WithinDuration(t, start.Add(test.want), deadline, time.Second)
		})
	}
}

func TestTransparentWrappersForwardAcquisitionBudget(t *testing.T) {
	t.Parallel()
	const login = 7 * time.Second
	federation := newFederation(auth.WithFederationAuthTimeout(login))
	for _, test := range []struct {
		name    string
		tokener auth.BearerTokener
	}{
		{name: "in-app sync", tokener: auth.NewInAppSyncTokener(federation)},
		{name: "name wrapper", tokener: auth.NewNameWrapper("name", federation)},
		{name: "instrumented", tokener: auth.NewInstrumentedBearerTokener(federation)},
		{name: "file cache", tokener: auth.NewFileCacheTokener(federation)},
		{name: "async file cache", tokener: auth.NewAsynchronouslyRenewableFileCacheTokener(federation)},
		{name: "cached bearer", tokener: auth.NewCachedBearerTokener(federation)},
		{name: "cached service", tokener: auth.NewCachedTokener(federation)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			provider, ok := test.tokener.(auth.AcquisitionBudgetProvider)
			require.True(t, ok, "wrapper does not implement the contract")
			budget, ok := provider.AcquisitionBudget()
			require.True(t, ok)
			require.Equal(t, login, budget)
		})
	}
}
