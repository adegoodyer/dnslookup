# dnslookup

[![Go Reference](https://pkg.go.dev/badge/github.com/adegoodyer/dnslookup.svg)](https://pkg.go.dev/github.com/adegoodyer/dnslookup)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

DNS Lookup CLI tool implemented in Go

## Install

```bash
go install github.com/adegoodyer/dnslookup/cmd/dnslookup@latest
```

## Usage
```bash
Usage: dnslookup <domain|IP>

Description:
This application performs DNS and reverse DNS lookups.
You can provide either a domain name or an IP address.

If a domain is provided, it will perform the following lookups:
 - A (IPv4) records
 - AAAA (IPv6) records
 - CNAME (Canonical Name) records
 - MX (Mail Exchange) records
 - NS (Name Server) records
 - TXT (Text) records
 - SOA (Start of Authority) record
 - WHOIS information for the domain

If an IP address is provided, it will perform reverse DNS lookup and display associated hostnames.

Example Usage:
  dnslookup example.com
  dnslookup 8.8.8.8
```

## Tags

- `latest`: Most recent stable build
- `x.y.z`: Specific version builds (e.g., `2.7.5`)
- `x.y`: Minor version builds (e.g., `2.7`)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
