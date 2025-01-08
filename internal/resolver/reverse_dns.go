package resolver

import (
	"fmt"
	"strings"
)

func PerformReverseLookup(resolver Resolver, ip string) {
	header := fmt.Sprintf("Reverse DNS Lookup Results for IP: %s", ip)
	fmt.Println("\n" + header)
	fmt.Println(strings.Repeat("=", len(header)))

	hostnames, err := resolver.LookupAddr(ip)
	if err != nil {
		fmt.Printf("Reverse DNS Lookup: Failed (%v)\n", err)
		return
	}

	fmt.Println("\nHostnames:")
	for _, name := range hostnames {
		fmt.Printf("- %s\n", name)
	}
}
