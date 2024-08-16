IN_PROGRESS = "is in progress ..."
DB_USER = "root"
DB_PASS = "pwd"
DB_HOST = "127.0.0.1"
DB_PORT = "3306"
DB_NAME = "acte"

## help: prints this help message
.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## mod: will pull all dependency
.PHONY: mod
mod:
	@echo "make mod ${IS_IN_PROGRESS}"
	@rm -rf ./vendor ./go.sum
	@go mod tidy
	@go mod vendor

## migrate-create-file: run create file migration scripts.
.PHONY: migrate-create-file
migrate-create-file:
	@goose -dir databases/migrations create ${FILENAME} sql
 

 ## setup: Set up database temporary for local environtment
.PHONY: setup
setup:
	@echo "make setup ${IS_IN_PROGRESS}"
	@docker-compose -f ./infrastructure/docker-compose.local.yml up -d
	@sleep 8

## down: Set down database temporary for local environtment
.PHONY: down
down: 
	@echo "make down ${IS_IN_PROGRESS}"
	@docker-compose -f ./infrastructure/docker-compose.local.yml down -t 1


.PHONY: migrate-up
migrate-up:
	@goose -dir databases/migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up  

PHONY: migrate-down
migrate-down:
	@goose -dir databases/migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" down