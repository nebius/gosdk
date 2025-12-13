package nid

import (
	"errors"
	"fmt"
	"regexp"
)

const maximumNidLength = 1024

var typeRegex = regexp.MustCompile(`^[a-z][a-z0-9]{2,49}$`)
var routingCodeRegex = regexp.MustCompile(`^[a-z][a-z0-9]{2}$`)
var weakIDRegex = regexp.MustCompile(`^[a-z0-9-]{1,71}[a-z0-9]$`)
var reservedRegex = regexp.MustCompile(`^([a-z-][a-z0-9-]{0,9})?$`)

var nidRegex = regexp.MustCompile(`^([a-z][a-z0-9]{2,49})-([a-z][a-z0-9]{2})(.+?)(?:--([a-z-][a-z0-9-]{0,9}))?$`)

//nolint:nonamedreturns,nakedret // for docs
func matchNid(s string) (nidType, nidRoutingCode, weakID, reserved string, err error) {
	if s == "" {
		err = errors.New("empty string")
		return
	}
	if len(s) > maximumNidLength {
		err = errors.New("maximum NID length exceeded")
		return
	}
	match := nidRegex.FindStringSubmatch(s)
	if match == nil {
		err = fmt.Errorf("nid %q does not match to %s", s, nidRegex)
		return
	}

	nidType = match[1]
	nidRoutingCode = match[2]
	weakID = match[3]
	reserved = match[4]

	return
}
