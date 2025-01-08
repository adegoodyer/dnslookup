# Tester Readme

## Overview
- MockResolver designed to simulate behaviour of Resolver interface
- each method in MockResolver has corresponding `Mock<MethodName>` field (e.g., `MockLookupAddr`)
- this field (a function) can be overridden to allow custom behaviour to be injected for specific method being tested

## Example In Code
```go
// mock_resolver.go

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
```

```go
// resolver_test.go

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
```

## Commands
```bash
# run all tests
go test ./...
go test ./... -v

# run tests for specific package
go test ./internal/resolver
go test ./internal/resolver -v

# run specific test
go test ./... -run TestLookupAddr
go test ./... -run TestLookupAddr -v

# show code coverage
go test ./... -cover

# generate coverage report
# (all test must pass beforehand for a complete report)
go test ./... -coverprofile=coverage.out

# show report in terminal
go tool cover -func=coverage.out

# show report in browser
go tool cover -html=coverage.out
```
