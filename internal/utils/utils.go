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

// prints usage instructions for application
func PrintUsage(appName string) {
	fmt.Printf("Usage: %s <domain|IP>\n", appName)
}
