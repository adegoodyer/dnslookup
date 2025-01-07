# Dev Readme

## Notes
- most of heavy lifting is done by `net` package
  - reverse lookup        net.LookupAddr(ip)
  - A and AAAA records    net.LookupIP(domain)
  - CNAME                 net.LookupCNAME(domain)
  - MX                    net.LookupMX(domain)
  - NS                    net.LookupNS(domain)
  - TXT                   net.LookupTXT(domain)
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

# install dependencies
go get k8s.io/client-go@latest
go get k8s.io/apimachinery@latest
go get github.com/spf13/cobra@latest
go get github.com/fatih/color

# tidy package dependencies
go mod tidy

# run
go run cmd/$PROJECT_NAME/main.go
```
