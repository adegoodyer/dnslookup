package lookups

import (
	"fmt"
	"net"
	"strings"
)

func PerformDNSLookups(domain string) {
	header := fmt.Sprintf("DNS Lookup Results for: %s", domain)
	fmt.Println("\n" + header)
	fmt.Println(strings.Repeat("=", len(header)))

	lookupA(domain)
	lookupAAAA(domain)
	lookupCNAME(domain)
	lookupMX(domain)
	lookupNS(domain)
	lookupTXT(domain)
}

func lookupA(domain string) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		fmt.Printf("A Record: Failed (%v)\n", err)
		return
	}

	fmt.Println("\nA Records (IPv4):")
	for _, ip := range ips {
		// check valid IPv4 address
		if ip.To4() != nil {
			fmt.Printf("- %s\n", ip)
		}
	}
}

func lookupAAAA(domain string) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		fmt.Printf("AAAA Record: Failed: (%v)\n", err)
		return
	}

	fmt.Println("\nAAAA Records (IPv6):")
	for _, ip := range ips {
		// check valid IPv6 address
		if ip.To16() != nil && ip.To4() == nil {
			fmt.Printf("- %s\n", ip)
		}
	}
}

func lookupCNAME(domain string) {
	cname, err := net.LookupCNAME(domain)

	if err != nil {
		fmt.Printf("CNAME Record: Failed (%v)\n", err)
		return
	}

	fmt.Println("\nCNAME Record:")
	fmt.Printf("- %s\n", cname)
}

func lookupMX(domain string) {
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		fmt.Printf("MX Record: Failed (%v)\n", err)
		return
	}

	fmt.Println("\nMX Records (Mail Servers):")
	for _, mx := range mxRecords {
		fmt.Printf("- %s (Priority: %d)\n", mx.Host, mx.Pref)
	}
}

func lookupNS(domain string) {
	nsRecords, err := net.LookupNS(domain)

	if err != nil {
		fmt.Printf("NS Record: Failed (%v)\n", err)
		return
	}

	fmt.Println("\nNS Records (Name Servers):")
	for _, ns := range nsRecords {
		fmt.Printf("- %s\n", ns.Host)
	}
}

func lookupTXT(domain string) {
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		fmt.Printf("TXT Record: Failed (%v)\n", err)
		return
	}

	fmt.Println("\nTXT Records:")
	for _, txt := range txtRecords {
		fmt.Printf("- %s\n", txt)
	}
}
