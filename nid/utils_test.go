package nid_test

import (
	"testing"

	"github.com/nebius/gosdk/nid"
)

func assertEmpty(actual nid.Reserved, reason string, t *testing.T) {
	if actual != nid.NoReserved {
		t.Errorf("Expected empty, but found %+v: %s", actual, reason)
	}
}

func assertEquals(expected, actual any, reason string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %+v, but found %+v: %s", expected, actual, reason)
	}
}

func assertNotEquals(expected, actual any, reason string, t *testing.T) {
	if expected == actual {
		t.Errorf("Expected not %+v, but found %+v: %s", expected, actual, reason)
	}
}

func do[T any](res T, err error) result[T] {
	return result[T]{
		Value: res,
		Error: err,
	}
}

func assertNoError[T any](res result[T], t *testing.T) T {
	if res.Error != nil {
		t.Errorf("Expected no error, but found %s", res.Error)
	}

	return res.Value
}

func assertNoErrorValue(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected no error, but found %s", err)
	}
}

func assertError[T any](res result[T], t *testing.T) T {
	if res.Error == nil {
		t.Errorf("Expected error")
	}

	return res.Value
}

func assertErrorValue(err error, reason string, t *testing.T) {
	if err == nil {
		t.Errorf("Expected error because %q, got nil", reason)
	}
}

type result[T any] struct {
	Value T
	Error error
}
