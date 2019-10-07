package AddrLookup

import (
	"testing"
	"net"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"fmt"
)

func TestAddrLookupSuccess(t *testing.T) {
	actualResult, _ := GetAddrHost("8.8.8.8")

	expectedResult, _ := net.LookupAddr("8.8.8.8")

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

func TestAddrLookupFailure(t *testing.T) {
	actualResult, _ := GetAddrHost("8.8.8.8")

	expectedResult, _ := net.LookupAddr("8.8.4.4")

	if !(cmp.Equal(actualResult, expectedResult)) {
		t.Fatal("IP does returns expected hostname when it should not")
	}
}