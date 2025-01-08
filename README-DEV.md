# Dev Readme

## Notes
- most of heavy lifting is done by `net` package
  - reverse lookup        net.LookupAddr(ip)
  - A and AAAA records    net.LookupIP(domain)
  - CNAME                 net.LookupCNAME(domain)
  - MX                    net.LookupMX(domain)
  - NS                    net.LookupNS(domain)
  - TXT                   net.LookupTXT(domain)
  - check valid IP        net.ParseIP(input)
- checking of valid IP addresses is interesting
  - converts to 4 or 16 byte representations - returns nil if not valid IP address
  - `if ip.To4() != nil`
  - `if ip.To16() != nil && ip.To4() == nil`

## Commands
```bash
# export envars
export PROJECT_NAME=dnslookup

# create gh repo
gh repo create adegoodyer/$PROJECT_NAME \
--description "" \
--add-readme \
-- private

# clone repo
gh repo clone adegoodyer/$PROJECT_NAME

# cd into repo
cd $PROJECT_NAME

# create go.mod
go mod init github.com/adegoodyer/$PROJECT_NAME

# setup project structure
mkdir -p bin cmd/$PROJECT_NAME internal && \
touch cmd/$PROJECT_NAME/main.go

# add .gitignore
# https://www.toptal.com/developers/gitignore/api/go

# get dependencies
go get github.com/stretchr/testify/mock
go get github.com/stretchr/testify/assert
go get github.com/miekg/dns
go get github.com/domainr/whois

# tidy package dependencies
go mod tidy

# run
go run cmd/$PROJECT_NAME/main.go

# git tags
# removes existing latest tag if it exists
# then creates new release version and latest tag
export PROJECT_VERSION=v0.0.2 && \
git tag -d latest 2>/dev/null && \
git push origin --delete latest 2>/dev/null || true && \
git tag -a $PROJECT_VERSION -m "Release version $PROJECT_VERSION" && \
git push origin $PROJECT_VERSION && \
git tag -a latest -m "Latest release" && \
git push origin latest

# build binary
go build -o bin/dnslookup ./cmd/dnslookup
```
