package utils

import (
	"net"
	"strings"
)

// check if input is valid IP address
func IsIP(input string) bool {
	return net.ParseIP(input) != nil
}

// check if input is valid domain name
func IsValidDomain(domain string) bool {
	// basic validation: no scheme, no spaces
	return !strings.Contains(domain, "://") && !strings.Contains(domain, " ")
}
