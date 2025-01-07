package resolver

import "net"

type Resolver interface {
	LookupAddr(ip string) ([]string, error)
	LookupIP(host string) ([]net.IP, error)
	LookupCNAME(host string) (string, error)
	LookupMX(host string) ([]*net.MX, error)
	LookupNS(host string) ([]*net.NS, error)
	LookupTXT(host string) ([]string, error)
}

// DefaultResolver uses Go's net package for lookups
type DefaultResolver struct{}

func (r *DefaultResolver) LookupAddr(ip string) ([]string, error) {
	return net.LookupAddr(ip)
}

func (r *DefaultResolver) LookupIP(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}

func (r *DefaultResolver) LookupCNAME(host string) (string, error) {
	return net.LookupCNAME(host)
}

func (r *DefaultResolver) LookupMX(host string) ([]*net.MX, error) {
	return net.LookupMX(host)
}

func (r *DefaultResolver) LookupNS(host string) ([]*net.NS, error) {
	return net.LookupNS(host)
}

func (r *DefaultResolver) LookupTXT(host string) ([]string, error) {
	return net.LookupTXT(host)
}
