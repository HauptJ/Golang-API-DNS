package HostLookup

import(
	"log"
	"net"
)

func GetHostRecord(host string) ([]string, error) {

	HostRecord, err := net.LookupHost(host)
	if err != nil {
		log.Printf("Problem getting A and / or AAAA records %v\n", err)
	}

	return HostRecord, err
}