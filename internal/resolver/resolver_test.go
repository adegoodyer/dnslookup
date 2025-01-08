package resolver

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupAddr(t *testing.T) {
	// creates new instance and pointer to it
	mockResolver := &MockResolver{
		// function assigned to field, overriding default behaviour for the test
		MockLookupAddr: func(ip string) ([]string, error) {
			return []string{"mock.hostname.com"}, nil
		},
	}

	// perform test
	hostnames, err := mockResolver.LookupAddr("127.0.0.1")
	assert.Nil(t, err)
	assert.Equal(t, []string{"mock.hostname.com"}, hostnames)
}

// pattern remains same for remaining tests..

func TestLookupIP(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupIP: func(host string) ([]net.IP, error) {
			return []net.IP{net.ParseIP("8.8.8.8")}, nil
		},
	}

	ips, err := mockResolver.LookupIP("example.com")
	assert.Nil(t, err)
	assert.Len(t, ips, 1)
	assert.Equal(t, "8.8.8.8", ips[0].String())
}

func TestLookupCNAME(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupCNAME: func(host string) (string, error) {
			return "mock.cname.com", nil
		},
	}

	cname, err := mockResolver.LookupCNAME("example.com")
	assert.Nil(t, err)
	assert.Equal(t, "mock.cname.com", cname)
}

func TestLookupMX(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupMX: func(host string) ([]*net.MX, error) {
			return []*net.MX{
				{Host: "mail.example.com", Pref: 10},
			}, nil
		},
	}

	mxRecords, err := mockResolver.LookupMX("example.com")
	assert.Nil(t, err)
	assert.Len(t, mxRecords, 1)
	assert.Equal(t, "mail.example.com", mxRecords[0].Host)
	assert.Equal(t, uint16(10), mxRecords[0].Pref)
}

func TestLookupNS(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupNS: func(host string) ([]*net.NS, error) {
			return []*net.NS{
				{Host: "ns.example.com"},
			}, nil
		},
	}

	nsRecords, err := mockResolver.LookupNS("example.com")
	assert.Nil(t, err)
	assert.Len(t, nsRecords, 1)
	assert.Equal(t, "ns.example.com", nsRecords[0].Host)
}

func TestLookupTXT(t *testing.T) {
	mockResolver := &MockResolver{
		MockLookupTXT: func(host string) ([]string, error) {
			return []string{"v=spf1 include:_spf.example.com ~all"}, nil
		},
	}

	txtRecords, err := mockResolver.LookupTXT("example.com")
	assert.Nil(t, err)
	assert.Len(t, txtRecords, 1)
	assert.Equal(t, "v=spf1 include:_spf.example.com ~all", txtRecords[0])
}

func TestLookupSOA(t *testing.T) {
	mockResolver := &MockResolver{}

	// Test SOA lookup
	soaRecords, err := mockResolver.LookupSOA("example.com")
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"Master Name Server: sns.dns.icann.org",
		"Responsible Email: noc.dns.icann.org",
		"Serial: 2020123456",
		"Refresh: 7200 seconds",
		"Retry: 3600 seconds",
		"Expire: 1209600 seconds",
		"Minimum TTL: 3600 seconds",
	}, soaRecords)
}

func TestLookupWHOIS(t *testing.T) {
	mockResolver := &MockResolver{}

	// Test WHOIS lookup
	whoisInfo, err := mockResolver.LookupWHOIS("example.com")
	assert.Nil(t, err)
	assert.Contains(t, whoisInfo, "Domain Name: EXAMPLE.COM")
	assert.Contains(t, whoisInfo, "Registrar WHOIS Server: whois.iana.org")
}
