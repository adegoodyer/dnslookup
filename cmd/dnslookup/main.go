package main

import (
	"os"

	"github.com/adegoodyer/dnslookup/internal/resolver"
	"github.com/adegoodyer/dnslookup/internal/utils"
)

func main() {
	// ensure domain/IP is provided
	if len(os.Args) < 2 {
		utils.PrintUsage(os.Args[0])
		os.Exit(1)
	}

	// sanitize input
	input := utils.CleanInput(os.Args[1])

	// check input is an IP address
	// perform relevant logic based on result
	if utils.IsIP(input) {
		resolver.PerformReverseLookup(&resolver.DefaultResolver{}, input)
	} else {
		resolver.PerformDNSLookups(&resolver.DefaultResolver{}, input)
	}
}
