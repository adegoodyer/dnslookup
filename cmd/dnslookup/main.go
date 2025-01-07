package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adegoodyer/dnslookup/internal/lookups"
	"github.com/adegoodyer/dnslookup/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <domain|IP>\n", os.Args[0])
		os.Exit(1)
	}

	input := strings.TrimSpace(os.Args[1])

	// validate input
	if utils.IsIP(input) {
		lookups.PerformReverseLookup(input)
	} else if utils.IsValidDomain(input) {
		lookups.PerformDNSLookups(input)
	} else {
		fmt.Println("Invalid input. Provide a valid IP or domain.")
		os.Exit(1)
	}
}
