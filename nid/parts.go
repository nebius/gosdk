package nid

import (
	"fmt"
	"strings"
)

type Type string
type RoutingCode string
type WeakID string
type Reserved string

const NoReserved = Reserved("")

// Deprecated: This constant only for backward compatibility of code.
// List of routing codes here: https://docs.nebius.dev/en/iam/about-iam/regions
const defaultRegion = RoutingCode("e0t")

func NewType(s string) (Type, error) {
	t := Type(s)

	if !t.IsValid() {
		return t, fmt.Errorf("type %q is invalid", s)
	}

	return t, nil
}

func NewRoutingCode(s string) (RoutingCode, error) {
	r := RoutingCode(s)

	if !r.IsValid() {
		return r, fmt.Errorf("routing code %q is invalid", s)
	}

	return r, nil
}

func NewWeakID(s string) (WeakID, error) {
	w := WeakID(s)

	if !w.IsValid() {
		return w, fmt.Errorf("weak id %q is invalid", s)
	}

	return w, nil
}

func NewReserved(s string) (Reserved, error) {
	r := Reserved(s)

	if !r.IsValid() {
		return r, fmt.Errorf("reserved %q is invalid", s)
	}

	return r, nil
}

func (t Type) IsValid() bool {
	return typeRegex.MatchString(string(t))
}

func (t Type) String() string {
	return string(t)
}

func (r RoutingCode) IsValid() bool {
	return routingCodeRegex.MatchString(string(r))
}

func (r RoutingCode) String() string {
	return string(r)
}

func (w WeakID) IsValid() bool {
	return weakIDRegex.MatchString(string(w)) && !strings.Contains(string(w), "--")
}

func (w WeakID) String() string {
	return string(w)
}

func (r Reserved) IsValid() bool {
	return reservedRegex.MatchString(string(r))
}
