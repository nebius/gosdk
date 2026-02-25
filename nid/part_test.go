package nid_test

import (
	"testing"

	"github.com/nebius/gosdk/nid"
)

func TestShouldSuccessfullyCreateNidType(t *testing.T) {
	t.Parallel()
	assertNoError(do(nid.NewType("abc")), t)
	assertNoError(do(nid.NewType("mytype")), t)
	assertNoError(do(nid.NewType("longtype50chars12345678901234567890123456789012345")), t)
}

func TestShouldSuccessfullyCreateNidRoutingCode(t *testing.T) {
	t.Parallel()
	assertNoError(do(nid.NewRoutingCode("e0t")), t)
	assertNoError(do(nid.NewRoutingCode("g0t")), t)
	assertNoError(do(nid.NewRoutingCode("xyz")), t)
}

func TestShouldSuccessfullyCreateNidWeakId(t *testing.T) {
	t.Parallel()
	assertNoError(do(nid.NewWeakID("abcde")), t)
	assertNoError(do(nid.NewWeakID("a7b1ccdde35fgghdii")), t)
	assertNoError(do(nid.NewWeakID("thisisuid123456789")), t)
	assertNoError(do(nid.NewWeakID("this-is-id-also")), t)
	assertNoError(do(nid.NewWeakID("this-is-idalso")), t)
	assertNoError(do(nid.NewWeakID("11111112222333444")), t)
	assertNoError(do(nid.NewWeakID("-sha256-5adefc5ff7bd508658fccef8b0ff217931b4392c59951f732f218780e376ceb4")), t)
	assertNoError(do(nid.NewWeakID("xxsha2565adefc5ff7bd508658fccef8b0ff217931b4392c59951f732f218780e376ceb4")), t)
}

func TestShouldSuccessfullyCreateNidReserved(t *testing.T) {
	t.Parallel()
	assertNoError(do(nid.NewReserved("-a-dsd-")), t)
	assertNoError(do(nid.NewReserved("--adsd-")), t)
	assertNoError(do(nid.NewReserved("---dsd-")), t)
	assertNoError(do(nid.NewReserved("----sd-")), t)
	assertNoError(do(nid.NewReserved("-----d-")), t)
	assertNoError(do(nid.NewReserved("a-dsd-")), t)
	assertNoError(do(nid.NewReserved("a--dsd-")), t)
	assertNoError(do(nid.NewReserved("abc")), t)
	assertNoError(do(nid.NewReserved("ab--c--")), t)
	assertNoError(do(nid.NewReserved("ab--c--d")), t)
	assertNoError(do(nid.NewReserved("a1b2c3d4e5")), t)
}

func TestShouldFailWhenNidTypeContainsHyphen(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewType("-")), t)
	assertError(do(nid.NewType("cool-type")), t)
	assertError(do(nid.NewType("co-ol-type")), t)
	assertError(do(nid.NewType("cooltype-")), t)
	assertError(do(nid.NewType("-cooltype")), t)
	assertError(do(nid.NewType("--")), t)
	assertError(do(nid.NewType("cool--type")), t)
	assertError(do(nid.NewType("co--ol--type")), t)
	assertError(do(nid.NewType("cooltype--")), t)
	assertError(do(nid.NewType("--cooltype")), t)
	assertError(do(nid.NewType("co-ol--type")), t)
	assertError(do(nid.NewType("co--ol-type")), t)
}

func TestShouldFailWhenNidRoutingCodeContainsHyphen(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewRoutingCode("-cf")), t)
	assertError(do(nid.NewRoutingCode("c-f")), t)
	assertError(do(nid.NewRoutingCode("cf-")), t)
	assertError(do(nid.NewRoutingCode("-c-")), t)
}

func TestShouldFailWhenNidWeakIdContainsDoubleHyphen(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewWeakID("--")), t)
	assertError(do(nid.NewWeakID("cool--weakid")), t)
	assertError(do(nid.NewWeakID("co--ol--weakid")), t)
	assertError(do(nid.NewWeakID("coolweakid--")), t)
	assertError(do(nid.NewWeakID("--coolweakid")), t)
	assertError(do(nid.NewWeakID("co-ol--weakid")), t)
	assertError(do(nid.NewWeakID("co--ol-weakid")), t)
}

func TestShouldFailWhenNidWeakIdContainsHyphenAtEnd(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewWeakID("aaaaaa-")), t)
	assertError(do(nid.NewWeakID("-aaaaa-")), t)
}

func TestShouldFailWhenNidTypeOrNidRoutingCodeOrNidReservedIsStartsFromDigit(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewType("1ytype")), t)
	assertError(do(nid.NewRoutingCode("1cf")), t)
	assertError(do(nid.NewReserved("1eserved")), t)
}

func TestShouldFailWhenPartContainsCharInWrongCase(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewType("mytYpe")), t)
	assertError(do(nid.NewRoutingCode("cCf")), t)
	assertError(do(nid.NewWeakID("yet-another-well-knOwn-id")), t)
	assertError(do(nid.NewReserved("reServed")), t)
}

func TestShouldFailWhenPartContainsBadChar(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewType("my-tяpe")), t)
	assertError(do(nid.NewRoutingCode("cяf")), t)
	assertError(do(nid.NewWeakID("yet-another-well-knяwn-id")), t)
	assertError(do(nid.NewReserved("resяrved")), t)
}

func TestShouldFailWhenPartIsTooLong(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewType("longtype51chars12345678901234567890123456789012345x")), t)
	assertError(do(nid.NewRoutingCode("cff4")), t)
	assertError(do(nid.NewWeakID(
		"longwkid98chars01234567890123456789012345678901234567890123456789012345678901234567890123456789xyz",
	)), t)
	assertError(do(nid.NewReserved("r11c012345x")), t)
}

func TestShouldFailWhenPartIsTooShort(t *testing.T) {
	t.Parallel()
	assertError(do(nid.NewType("ab")), t)
	assertError(do(nid.NewRoutingCode("")), t)
	assertError(do(nid.NewRoutingCode("c")), t)
	assertError(do(nid.NewRoutingCode("cf")), t)
	assertError(do(nid.NewWeakID("a")), t)
	// assertError(do(nid.NewReserved("")), t) // NOTE(anton): empty Reserved === no Reserved
}
