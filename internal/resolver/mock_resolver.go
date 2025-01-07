package resolver

import "net"

// MockResolver is a struct used to simulate resolver behavior for testing purposes.
type MockResolver struct {
	MockLookupAddr  func(string) ([]string, error)
	MockLookupIP    func(string) ([]net.IP, error)
	MockLookupCNAME func(string) (string, error)
	MockLookupMX    func(string) ([]*net.MX, error)
	MockLookupNS    func(string) ([]*net.NS, error)
	MockLookupTXT   func(string) ([]string, error)
}

// LookupAddr simulates reverse DNS lookup for testing.
func (m *MockResolver) LookupAddr(ip string) ([]string, error) {
	if m.MockLookupAddr != nil {
		return m.MockLookupAddr(ip)
	}
	return nil, nil
}

// LookupIP simulates DNS lookup for A or AAAA records for testing.
func (m *MockResolver) LookupIP(host string) ([]net.IP, error) {
	if m.MockLookupIP != nil {
		return m.MockLookupIP(host)
	}
	return nil, nil
}

// LookupCNAME simulates DNS lookup for CNAME records for testing.
func (m *MockResolver) LookupCNAME(host string) (string, error) {
	if m.MockLookupCNAME != nil {
		return m.MockLookupCNAME(host)
	}
	return "", nil
}

// LookupMX simulates DNS lookup for MX records for testing.
func (m *MockResolver) LookupMX(host string) ([]*net.MX, error) {
	if m.MockLookupMX != nil {
		return m.MockLookupMX(host)
	}
	return nil, nil
}

// LookupNS simulates DNS lookup for NS records for testing.
func (m *MockResolver) LookupNS(host string) ([]*net.NS, error) {
	if m.MockLookupNS != nil {
		return m.MockLookupNS(host)
	}
	return nil, nil
}

// LookupTXT simulates DNS lookup for TXT records for testing.
func (m *MockResolver) LookupTXT(host string) ([]string, error) {
	if m.MockLookupTXT != nil {
		return m.MockLookupTXT(host)
	}
	return nil, nil
}
