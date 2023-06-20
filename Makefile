run-lint:
	golangci-lint run ./... -v

generate:
	go generate ./...