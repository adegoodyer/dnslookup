package resolver

import "net"

// MockResolver is struct used to simulate resolver behaviour for testing purposes
type MockResolver struct {
	MockLookupAddr  func(string) ([]string, error)
	MockLookupIP    func(string) ([]net.IP, error)
	MockLookupCNAME func(string) (string, error)
	MockLookupMX    func(string) ([]*net.MX, error)
	MockLookupNS    func(string) ([]*net.NS, error)
	MockLookupTXT   func(string) ([]string, error)
}

// simulates reverse DNS lookup for testing
func (m *MockResolver) LookupAddr(ip string) ([]string, error) {
	// checks if field is set
	if m.MockLookupAddr != nil {
		//if it is, then calls function provided by overridden method
		return m.MockLookupAddr(ip)
	}
	// if not, returns default values of nil, nil
	return nil, nil
}

// pattern remains same for remaining methods..

// simulates DNS lookup for A or AAAA records for testing
func (m *MockResolver) LookupIP(host string) ([]net.IP, error) {
	if m.MockLookupIP != nil {
		return m.MockLookupIP(host)
	}
	return nil, nil
}

// simulates DNS lookup for CNAME records for testing
func (m *MockResolver) LookupCNAME(host string) (string, error) {
	if m.MockLookupCNAME != nil {
		return m.MockLookupCNAME(host)
	}
	return "", nil
}

// simulates DNS lookup for MX records for testing
func (m *MockResolver) LookupMX(host string) ([]*net.MX, error) {
	if m.MockLookupMX != nil {
		return m.MockLookupMX(host)
	}
	return nil, nil
}

// simulates DNS lookup for NS records for testing
func (m *MockResolver) LookupNS(host string) ([]*net.NS, error) {
	if m.MockLookupNS != nil {
		return m.MockLookupNS(host)
	}
	return nil, nil
}

// simulates DNS lookup for TXT records for testing
func (m *MockResolver) LookupTXT(host string) ([]string, error) {
	if m.MockLookupTXT != nil {
		return m.MockLookupTXT(host)
	}
	return nil, nil
}

func (r *MockResolver) LookupSOA(host string) ([]string, error) {
	return []string{
		"Master Name Server: sns.dns.icann.org",
		"Responsible Email: noc.dns.icann.org",
		"Serial: 2020123456",
		"Refresh: 7200 seconds",
		"Retry: 3600 seconds",
		"Expire: 1209600 seconds",
		"Minimum TTL: 3600 seconds",
	}, nil
}

func (r *MockResolver) LookupWHOIS(domain string) (string, error) {
	return "Domain Name: EXAMPLE.COM\nRegistry Domain ID: D123456789-COM\nRegistrar WHOIS Server: whois.iana.org", nil
}
