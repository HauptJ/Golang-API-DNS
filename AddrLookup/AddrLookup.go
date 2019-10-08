package AddrLookup

import (
	"log"
	"net"
)

func GetAddrHost(hostIP string) ([]string, error) {
	
	HostNames, err := net.LookupAddr(hostIP)
	if err != nil {
        log.Printf("Problem getting MX record %v\n", err)
	}
	
	return HostNames, err
}