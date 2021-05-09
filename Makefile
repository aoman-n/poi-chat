.PHONY: in-api migrate-create migrate-up migrate-down

EXEC_API=docker-compose exec api

in-api:
	$(EXEC_API) /bin/bash

migrate-create:
	$(EXEC_API) goose create ${FILENAME} sql

migrate-up:
	$(EXEC_API) goose up

migrate-down:
	$(EXEC_API) goose down
