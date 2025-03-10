APP_NAME = server
GOOSE_DBSTRING = $(STR_MYSQL)"root:cmmtpnx1@tcp(127.0.0.1:3307)/bookinggo"
GOOSE_MIGRATION_DIR = sql/schema
GOOSE_DRIVER = mysql

docker_build:
	docker compose up -d --build 
	docker compose ps 

dev: 
	go run ./cmd/${APP_NAME}/main.go

run:
	docker compose up -d && go run ./cmd/${APP_NAME}

kill:
	docker compose kill

docker_up: 
	docker compose up -d

docker_down: 
	docker compose down

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@cmd /C "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) up"

downse:
	@cmd /C "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) down"

resetse:
	@cmd /C "set GOOSE_DRIVER=$(GOOSE_DRIVER)&& set GOOSE_DBSTRING=$(GOOSE_DBSTRING)&& goose -dir=$(GOOSE_MIGRATION_DIR) reset"

sqlgen:
	@powershell -Command "docker run --rm -v \"$${PWD}:/src\" -w /src sqlc/sqlc generate"

wiregen:
	cd internal\wire && wire

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev run downse upse resetse docker_up docker_down wiregen

.PHONY: air