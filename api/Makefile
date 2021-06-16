.PHONY: gen gqlgen fmt lint

gen:
	- go generate ./...

gqlgen:
	- gqlgen generate

fmt:
	gofmt -w ./src/...

lint:
	golangci-lint run --tests

test:
	GO_ENV=test \
	DB_NAME=poi_test \
	REDIS_DB=1 \
	LOG_LEVEL=info \
	go test -v ./src/repository/...
