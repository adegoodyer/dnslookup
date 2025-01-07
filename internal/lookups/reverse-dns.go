package lookups

import (
	"fmt"
	"net"
	"strings"
)

func PerformReverseLookup(ip string) {
	header := fmt.Sprintf("Reverse DNS Lookup Results for IP: %s", ip)
	fmt.Println("\n" + header)
	fmt.Println(strings.Repeat("=", len(header)))

	hostnames, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Printf("Reverse DNS Lookup: Failed (%v)\n", err)
	} else {
		fmt.Println("\nHostnames:")
		for _, name := range hostnames {
			fmt.Printf("- %s\n", name)
		}
	}
}
