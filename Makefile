clean:
	rm -rf tplink2mqtt.*
lint:
	golangci-lint run ./internal/... ./cmd/... ./pkg/...
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ldap.linux_amd64 ./cmd/ldap
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ldap.darwin_amd64 ./cmd/ldap
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ldap.windows_amd64.exe ./cmd/ldap

docker:
	docker build -f ./cmd/imap/Dockerfile -t shauncampbell/ldap:local .