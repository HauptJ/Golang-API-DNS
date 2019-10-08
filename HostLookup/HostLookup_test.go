package HostLookup

import (
	"testing"
	"net"
)

func Equal(a []string, b []string) (bool) {
	if (len(a) != len(b)) {
		return false
	}
	
	for i, host := range a {
		if (host != b[i]) {
			return false
		}
	}
	return true
}

func TestHostLookupSuccess(t *testing.T) {
	actualResult, _ := GetHostRecord("localhost")

	expectedResult, _ := net.LookupHost("localhost")

	if !Equal(actualResult, expectedResult) {
		t.Fatalf("hostname did not return expected IP")
	}
}