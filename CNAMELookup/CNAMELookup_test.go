package CNAMELookup

import (
	"testing"
	"net"
)

func TestCNAMELookupSuccess(t *testing.T) {
	actualResult, _ := GetCNAMERecord("research.swtch.com")

	expectedResult, _ := net.LookupCNAME("research.swtch.com")

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}