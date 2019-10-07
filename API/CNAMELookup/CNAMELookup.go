package CNAMELookup

import(
	"log"
	"net"
)

func GetCNAMERecord(host string) (string, error) {

	CNAMERecord, err := net.LookupCNAME(host)
	if err != nil {
		log.Printf("Problem getting CNAME record %v\n", err)
	}

	return CNAMERecord, err
}