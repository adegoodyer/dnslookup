package resolver

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test mock resolver's LookupAddr method
func TestMockResolver_LookupAddr(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupAddr: func(ip string) ([]string, error) {
			if ip == "127.0.0.1" {
				return []string{"mock.hostname.com"}, nil
			}
			return nil, nil
		},
	}

	// Test with a valid IP
	hostnames, err := mockResolver.LookupAddr("127.0.0.1")
	assert.Nil(t, err)
	assert.Equal(t, []string{"mock.hostname.com"}, hostnames)

	// Test with an invalid IP
	hostnames, err = mockResolver.LookupAddr("192.168.1.1")
	assert.Nil(t, err)
	assert.Empty(t, hostnames)
}

// Test mock resolver's LookupIP method
func TestMockResolver_LookupIP(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupIP: func(host string) ([]net.IP, error) {
			if host == "example.com" {
				return []net.IP{net.ParseIP("8.8.8.8")}, nil
			}
			return nil, nil
		},
	}

	// Test with a valid domain
	ips, err := mockResolver.LookupIP("example.com")
	assert.Nil(t, err)
	assert.Equal(t, net.ParseIP("8.8.8.8"), ips[0])

	// Test with an invalid domain
	ips, err = mockResolver.LookupIP("nonexistentdomain.com")
	assert.Nil(t, err)
	assert.Empty(t, ips)
}

// Test mock resolver's LookupCNAME method
func TestMockResolver_LookupCNAME(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupCNAME: func(host string) (string, error) {
			if host == "example.com" {
				return "cname.example.com", nil
			}
			return "", nil
		},
	}

	// Test with a valid domain
	cname, err := mockResolver.LookupCNAME("example.com")
	assert.Nil(t, err)
	assert.Equal(t, "cname.example.com", cname)

	// Test with an invalid domain
	cname, err = mockResolver.LookupCNAME("nonexistentdomain.com")
	assert.Nil(t, err)
	assert.Empty(t, cname)
}

// Test mock resolver's LookupMX method
func TestMockResolver_LookupMX(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupMX: func(host string) ([]*net.MX, error) {
			if host == "example.com" {
				return []*net.MX{
					{Host: "mail.example.com", Pref: 10},
				}, nil
			}
			return nil, nil
		},
	}

	// Test with a valid domain
	mxRecords, err := mockResolver.LookupMX("example.com")
	assert.Nil(t, err)
	assert.Equal(t, "mail.example.com", mxRecords[0].Host)

	// Test with an invalid domain
	mxRecords, err = mockResolver.LookupMX("nonexistentdomain.com")
	assert.Nil(t, err)
	assert.Empty(t, mxRecords)
}

// Test mock resolver's LookupNS method
func TestMockResolver_LookupNS(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupNS: func(host string) ([]*net.NS, error) {
			if host == "example.com" {
				return []*net.NS{
					{Host: "ns1.example.com"},
				}, nil
			}
			return nil, nil
		},
	}

	// Test with a valid domain
	nsRecords, err := mockResolver.LookupNS("example.com")
	assert.Nil(t, err)
	assert.Equal(t, "ns1.example.com", nsRecords[0].Host)

	// Test with an invalid domain
	nsRecords, err = mockResolver.LookupNS("nonexistentdomain.com")
	assert.Nil(t, err)
	assert.Empty(t, nsRecords)
}

// Test mock resolver's LookupTXT method
func TestMockResolver_LookupTXT(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupTXT: func(host string) ([]string, error) {
			if host == "example.com" {
				return []string{"v=spf1 include:_spf.example.com ~all"}, nil
			}
			return nil, nil
		},
	}

	// Test with a valid domain
	txtRecords, err := mockResolver.LookupTXT("example.com")
	assert.Nil(t, err)
	assert.Equal(t, "v=spf1 include:_spf.example.com ~all", txtRecords[0])

	// Test with an invalid domain
	txtRecords, err = mockResolver.LookupTXT("nonexistentdomain.com")
	assert.Nil(t, err)
	assert.Empty(t, txtRecords)
}
