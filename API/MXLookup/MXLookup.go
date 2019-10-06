package MXLookup

import (
    "log"
    "net"
)

func GetMXRecord(host string) ([]*net.MX, error) {
    
    MXrecord, err := net.LookupMX(host)
    if err != nil {
        log.Printf("Problem getting MX record %v\n", err)
    }

    return MXrecord, err
}