.PHONY: mockgen-repo gqlgen fmt lint test-repo

mockgen-domain:
	- go generate ./src/domain/.../.

gqlgen:
	- gqlgen generate

fmt:
	gofmt -w ./src/...

lint:
	golangci-lint run --tests

test-repo:
	GO_ENV=test \
	DB_NAME=poi_test \
	REDIS_DB=1 \
	LOG_LEVEL=info \
	go test -v ./src/repository/...
