.PHONY: gen gqlgen fmt

gen:
	- go generate ./...

gqlgen:
	- gqlgen generate

fmt:
	- gofmt -w ./src/**/*.go
