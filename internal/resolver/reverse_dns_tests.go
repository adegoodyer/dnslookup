package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformReverseLookup(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupAddr: func(ip string) ([]string, error) {
			// simulating mock behaviour for reverse lookup
			return []string{"mock.hostname.com"}, nil
		},
	}

	// capture output to validate expected behaviour
	var result []string
	mockResolver.MockLookupAddr = func(ip string) ([]string, error) {
		return []string{"mock.hostname.com"}, nil
	}

	// invoke PerformReverseLookup with mock resolver
	PerformReverseLookup(mockResolver, "127.0.0.1")

	// assert expected result
	assert.Equal(t, []string{"mock.hostname.com"}, result)
}
