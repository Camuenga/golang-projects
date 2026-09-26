package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	domainName := bufio.NewScanner(os.Stdin)

	fmt.Printf("Checking email.................. \n\n")
	fmt.Printf("Domain name example.com: ")
	for domainName.Scan() {
		verifyDomain(domainName.Text())

		if err := domainName.Err(); err != nil {
			log.Fatal("Error %v", err)
		}
	}

}

func verifyDomain(domainName string) {

	var spfRecord, dmarcRecord string
	var hasMX, hasSPF, hasDMARC bool

	if mxRecord, err := net.LookupMX(domainName); err != nil {
		log.Printf("MX: %v ", err)
	} else {
		if len(mxRecord) > 0 {
			hasMX = true
		}
	}

	if spfRecordTXT, err := net.LookupTXT("spf1." + domainName); err != nil {
		log.Printf("ERROR SPF: %v", err)
	} else {
		for _, record := range spfRecordTXT {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
			}
		}
	}

	if dmarcRecordTXT, err := net.LookupTXT("_dmarc." + domainName); err != nil {
		log.Printf("ERROR DMARC: %v", err)
	} else {
		for _, record := range dmarcRecordTXT {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record

			}
		}
	}

	fmt.Printf("Domain: %v, \nMX: %v, \nSPF: %v, \nSPF-RECORD: %v, \nDMARC: %v, \nDMARC-RECORD:  %v \n",
		domainName, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	fmt.Printf("\nDomain name example.com: ")

}
