package nid_test

import (
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/nebius/gosdk/nid"
)

func TestShouldParseUIDBasedNID(t *testing.T) {
	t.Parallel()
	assertParse("folder-e0taccdde35fgghdii", "folder", "e0t", "accdde35fgghdii", t)
	assertParse("folder-e0taccdde35fgghdii--abcde", "folder", "e0t", "accdde35fgghdii", t)
	assertParse("longtype50chars12345678901234567890123456789012345-e0taccdde35fgghdii--a1b2c3d4e5", "longtype50chars12345678901234567890123456789012345", "e0t", "accdde35fgghdii", t)
	assertParse("stp-e0taccdde35fgghdii", "stp", "e0t", "accdde35fgghdii", t)
}

func TestShouldParseWellKnownIDBasedNID(t *testing.T) {
	t.Parallel()
	assertParse("folder-e0t-my-cool-wkid", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0t-mycoolwkid", "folder", "e0t", "-mycoolwkid", t)
	assertParse("folder-e0t-my-cool-wkid--abcde", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0t-mycoolwkid--abcde", "folder", "e0t", "-mycoolwkid", t)
	assertParse("longtype50chars12345678901234567890123456789012345-e0t-sha256-5adefc5ff7bd508658fccef8b0ff217931b4392c59951f732f218780e376ceb4--a1b2c3d4e5", "longtype50chars12345678901234567890123456789012345", "e0t", "-sha256-5adefc5ff7bd508658fccef8b0ff217931b4392c59951f732f218780e376ceb4", t)
	assertParse("stp-e0t-swkid", "stp", "e0t", "-swkid", t)
}

func TestShouldParseNIDWithoutHyphensInPart(t *testing.T) {
	t.Parallel()
	assertParse("folder-e0tmycoolwkid", "folder", "e0t", "mycoolwkid", t)
	assertParse("folder-e0taccdde35fgghdii", "folder", "e0t", "accdde35fgghdii", t)
	assertParse("folder-e0tmycoolwkid--rsrvd", "folder", "e0t", "mycoolwkid", t)
}

func TestShouldParseNIDWithHyphensInPart(t *testing.T) {
	t.Parallel()
	assertParse("folder-e0t-my-cool-wkid", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0t-my-cool-wkid--rsrvd", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0t-my-cool-wkid--rs-rvd", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0t-my-cool-wkid--rs--rvd", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0t-mycoolwkid--rsrvd", "folder", "e0t", "-mycoolwkid", t)
	assertParse("folder-cf4-mycoolwkid--rs-rvd", "folder", "cf4", "-mycoolwkid", t)
	assertParse("folder-e0t-mycoolwkid--rs--rvd", "folder", "e0t", "-mycoolwkid", t)
	assertParse("folder-e0t-my-cool-wkid------", "folder", "e0t", "-my-cool-wkid", t)
	assertParse("folder-e0taccdde35fgghdii--rsrvd", "folder", "e0t", "accdde35fgghdii", t)
	assertParse("folder-e0taccdde35fgghdii--rs-rvd", "folder", "e0t", "accdde35fgghdii", t)
	assertParse("folder-e0taccdde35fgghdii--rs--rvd", "folder", "e0t", "accdde35fgghdii", t)
	assertParse("folder-cf4accdde35fgghdii------", "folder", "cf4", "accdde35fgghdii", t)
	assertParse("folder-e0taccdde--x5fgghdii", "folder", "e0t", "accdde", t)
	assertParse("folder-e0taccdde--x5fg--dii", "folder", "e0t", "accdde", t)
	assertParse("folder-e0t-ccdde--x5fg--dii", "folder", "e0t", "-ccdde", t)
	assertParse("folder-e0taccdde--x5-g--dii", "folder", "e0t", "accdde", t)
}

func TestShouldReturnErrorOnParseNIDWithInvalidFormat(t *testing.T) {
	t.Parallel()
	assertParseFailed("weakid", "no hyphens", t)
	assertParseFailed("1older-a7b1ccdde35fgghdii", "first char of type is digit", t)
	assertParseFailed("folder-17baccdde35fgghdii", "first char of routing code is digit", t)
	assertParseFailed("folder-17b1ccdde35fgghdii--1eserved", "first char of reserved is digit", t)
	assertParseFailed("foяder-a7b1ccdde35fgghdii", "bad char in type", t)
	assertParseFailed("folder-aяbaccdde35fgghdii", "bad char in routing code", t)
	assertParseFailed("folder-a7baяcdde35fgghdii", "bad char in weak-id", t)
	assertParseFailed("folder-e0taяcdde35fgghdii", "bad char in weak-id", t)
	assertParseFailed("folder-a7b1ccdde35fgghdii--resяrved", "bad char in reserved", t)
	assertParseFailed("folder-a7b1ccdde35fgghdii--", "empty reserved", t)
}

func TestShouldReturnErrorOnParseNIDWithWrongCase(t *testing.T) {
	t.Parallel()
	assertParseFailed("foLder-a7b1ccdde35fgghdii", "wrong case in type", t)
	assertParseFailed("folder-e0T1ccdde35fgghdii", "wrong case in routing code", t)
	assertParseFailed("folder-a7b1Ccdde35fgghdii", "wrong case in weak-id", t)
	assertParseFailed("folder-a7b1ccdde35fgghdii--rEserved", "wrong case in reserved", t)
}

func TestShouldReturnErrorOnParseNIDWithInvalidLengths(t *testing.T) {
	t.Parallel()
	assertParseFailed("longtype51chars12345678901234567890123456789012345x-a7b1ccdde35fgghdii", "type is too long", t)
	assertParseFailed("folder-long-well-known-id-98-chars-1234567890123456789012345678901234567890123456789012345678901234567890", "weak-id is too long", t)
	assertParseFailed("folder-a7b1ccdde35fgghdii--reserved123", "reserved is too long", t)
	assertParseFailed("ab-a7b1ccdde35fgghdii", "type is too short", t)
	assertParseFailed("folder-acd", "weak-id is too short", t)
	assertParseFailed("folder-a7b1ccdde35fgghdii--", "re is too short", t)
}

func assertParse(idStr, expectedType, expectedRoutingCode, expectedWeakID string, t *testing.T) {
	id := assertNoError(do(nid.Parse(idStr)), t)

	assertNoErrorValue(id.Validate(), t)

	assertValueScan(id, t)

	assertEquals(idStr, id.String(), "should turn to string", t)
	assertEquals(idStr, id.Serialize(), "should reserialize", t)

	assertEquals(expectedWeakID, string(id.WeakID), "nid.weakId", t)
	assertEquals(expectedType, string(id.Type), "nid.idType", t)
	assertEquals(expectedRoutingCode, string(id.RoutingCode), "nid.RoutingCode", t)
}

func assertParseFailed(idStr, reason string, t *testing.T) {
	id := assertError(do(nid.Parse(idStr)), t)

	assertErrorValue(id.Validate(), reason, t)
}

func assertValueScan(origID nid.Nid, t *testing.T) {
	value := assertNoError(do(origID.Value()), t)

	scanedID := new(nid.Nid)
	_ = assertNoError(do(scanedID, scanedID.Scan(value)), t)

	assertEquals(origID, *scanedID, "should be equal", t)
}

func TestNidYAMLSerialization(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		nid      nid.Nid
		expected string
	}{
		{
			name:     "Valid Nid",
			nid:      nid.Nid{Type: "usr", RoutingCode: "e0t", WeakID: "abc123", Reserved: nid.NoReserved},
			expected: "usr-e0tabc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			yamlData, err := yaml.Marshal(tt.nid)
			if err != nil {
				t.Fatalf("Failed to marshal Nid: %v", err)
			}

			var result string
			err = yaml.Unmarshal(yamlData, &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal YAML: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}

func TestNidYAMLDeserialization(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		yaml     string
		expected nid.Nid
		wantErr  bool
	}{
		{
			name:     "Valid Nid",
			yaml:     "usr-e0tabc123",
			expected: nid.Nid{Type: "usr", RoutingCode: "e0t", WeakID: "abc123", Reserved: nid.NoReserved},
			wantErr:  false,
		},
		{
			name:    "Invalid Nid",
			yaml:    "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			yamlData := []byte(tt.yaml)
			var result nid.Nid
			err := yaml.Unmarshal(yamlData, &result)

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestNidInStruct(t *testing.T) {
	t.Parallel()
	type TestStruct struct {
		Tenant nid.Nid `yaml:"tenant"`
	}

	original := TestStruct{
		Tenant: nid.Nid{Type: "org", RoutingCode: "e0t", WeakID: "abc123", Reserved: nid.NoReserved},
	}

	yamlData, err := yaml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal struct: %v", err)
	}

	var decoded TestStruct
	err = yaml.Unmarshal(yamlData, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal struct: %v", err)
	}

	if decoded.Tenant != original.Tenant {
		t.Errorf("Expected Tenant %v, but got %v", original.Tenant, decoded.Tenant)
	}
}
