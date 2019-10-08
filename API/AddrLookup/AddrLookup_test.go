package AddrLookup

import (
	"testing"
	"net"
	"fmt"
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

func TestAddrLookupSuccess(t *testing.T) {
	actualResult, _ := GetAddrHost("1.1.1.1")

	expectedResult, _ := net.LookupAddr("1.1.1.1")

	if !Equal(actualResult, expectedResult) {
		t.Fatal("IP does not return expected hostname")
	}
}