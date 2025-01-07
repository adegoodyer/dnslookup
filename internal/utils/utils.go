package utils

import (
	"fmt"
	"net"
	"strings"
)

// trims whitespace and normalizes input
func CleanInput(input string) string {
	return strings.TrimSpace(input)
}

// checks if given string is valid IP address
func IsIP(input string) bool {
	return net.ParseIP(input) != nil
}

// prints usage instructions for the application
func PrintUsage(appName string) {
	fmt.Printf("Usage: %s <domain|IP>\n\n", appName)
	fmt.Println("Description:")
	fmt.Println("This application performs DNS and reverse DNS lookups.")
	fmt.Println("You can provide either a domain name or an IP address.")
	fmt.Println("\nIf a domain is provided, it will perform the following lookups:")
	fmt.Println(" - A (IPv4) records")
	fmt.Println(" - AAAA (IPv6) records")
	fmt.Println(" - CNAME (Canonical Name) records")
	fmt.Println(" - MX (Mail Exchange) records")
	fmt.Println(" - NS (Name Server) records")
	fmt.Println(" - TXT (Text) records")
	fmt.Println(" - SOA (Start of Authority) record")
	fmt.Println(" - WHOIS information for the domain")
	fmt.Println("\nIf an IP address is provided, it will perform reverse DNS lookup and display associated hostnames.")
	fmt.Println("\nExample Usage:")
	fmt.Printf("  %s example.com\n", appName)
	fmt.Printf("  %s 8.8.8.8\n", appName)
	fmt.Println("\nFor more information, refer to the documentation.")
}
