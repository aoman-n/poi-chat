.PHONY: gen gqlgen fmt lint

gen:
	- go generate ./...

gqlgen:
	- gqlgen generate

fmt:
	- gofmt -w ./src/**/*.go

lint:
	golangci-lint run --tests
