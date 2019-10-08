package MXLookup

import (
	"testing"
	"net"
	"github.com/google/go-cmp/cmp"
)

func TestMXLookupSuccess(t *testing.T) {
	actualResult, _ := GetMXRecord("hauptj.com")

	expectedResult, _ := net.LookupMX("hauptj.com")

	if cmp.Equal(actualResult, expectedResult) {
		t.Fatalf("Records are not equal")
	}
}

func TestMXLookupFailure(t *testing.T) {
	actualResult, _ := GetMXRecord("net.hauptj.com")

	expectedResult, _ := net.LookupMX("net.hauptj.com")

	if !(cmp.Equal(actualResult, expectedResult)) {
		t.Fatalf("Records are not equal")
	}
}

