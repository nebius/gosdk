package auth

import "time"

// AcquisitionBudgetProvider is an optional contract of a [BearerTokener]. It
// reports how long one [BearerTokener.BearerToken] call through this tokener
// can take. A cache that bounds a shared acquisition asks only its immediate
// tokener for the budget. It does not look through the chain.
//
// A tokener that owns an operation timeout reports that timeout, for example
// the login timeout of [FederationTokener]. A tokener that adds an operation to
// its wrapped tokener reports the composed budget, for example
// [ExchangeImpersonatedBearerTokener]. A transparent wrapper forwards the
// budget of its wrapped tokener. A tokener without a budget does not implement
// the contract, and the cache uses its default.
//
// The second result is false when the tokener has no budget. A non-positive
// budget means that the operation fails at once; callers must not extend it.
type AcquisitionBudgetProvider interface {
	AcquisitionBudget() (time.Duration, bool)
}

// forwardAcquisitionBudget returns the budget of the wrapped tokener of a
// transparent wrapper. It checks only that tokener.
func forwardAcquisitionBudget(tokener BearerTokener) (time.Duration, bool) {
	if provider, ok := tokener.(AcquisitionBudgetProvider); ok {
		return provider.AcquisitionBudget()
	}
	return 0, false
}
