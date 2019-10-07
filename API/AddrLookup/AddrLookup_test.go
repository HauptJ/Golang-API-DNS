package AddrLookup

import (
	"testing"
	"net"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func TestAddrLookupSuccess(t *testing.T) {
	actualResult, _ := GetAddrHost("1.1.1.1")

	expectedResult, _ := net.LookupAddr("1.1.1.1")

	if cmp.Equal(actualResult, expectedResult) {
		for _, host := range actualResult {
			fmt.Printf("%s\n", host)
		}

		for _, host := range expectedResult {
			fmt.Printf("%s\n", host)
		}
		t.Fatal("IP does not return expected hostname")
	}
}