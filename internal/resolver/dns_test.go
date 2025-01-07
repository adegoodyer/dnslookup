package resolver

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformDNSLookups(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupIP: func(host string) ([]net.IP, error) {
			return []net.IP{net.ParseIP("8.8.8.8")}, nil
		},
	}

	// Test DNS lookup
	var result string
	mockResolver.MockLookupIP = func(host string) ([]net.IP, error) {
		result = "8.8.8.8"
		return []net.IP{net.ParseIP(result)}, nil
	}

	PerformDNSLookups(mockResolver, "example.com")
	assert.Equal(t, "8.8.8.8", result)
}
