package MXLookup

import (
	"testing"
	"net"
)

func TestMXLookupSuccess(t *testing.T) {
	actualResult, err := GetMXRecord("hauptj.com")
	if err != nil {
		panic(err)
	}
	expectedResult, err := net.LookupMX("hauptj.com")
	if err != nil {
		panic(err)
	}

	if &actualResult != &expectedResult {
		t.Fatalf("Records are not equal")
	}
}

