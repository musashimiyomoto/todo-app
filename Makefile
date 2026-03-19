include .env
export

export PROJECT_ROOT := $(shell pwd)

run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run cmd/app/main.go

up:
	@docker compose up --build

down:
	@docker compose down

cleanup:
	@read -p "Clean all data? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down && \
		sudo rm -rf out/pgdata && \
		echo "Clean success!"; \
	else \
		echo "Clean canceled!"; \
	fi

migrate-create:
	@if [ -z "${seq}" ]; then \
		echo "Missing value seq. Example make migrate-create seq=init"; \
		exit 1; \
	fi;
	@docker compose run --rm migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "${seq}"

migrate-action:
	@if [ -z "${action}" ]; then \
		echo "Missing value action. Example make migrate-action action=up"; \
		exit 1; \
	fi;
	@docker compose run --rm migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable \
		"${action}"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down
