include .env
export

export PROJECT_ROOT := $(shell pwd)

run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go run ${PROJECT_ROOT}/cmd/app/main.go

up:
	@docker compose up --build

down:
	@docker compose down

cleanup:
	@read -p "Clean all postgres data? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Clean postgres data success!"; \
	else \
		echo "Clean postgres data canceled!"; \
	fi

	@read -p "Clean all logs? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		sudo rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Clean logs success!"; \
	else \
		echo "Clean logs canceled!"; \
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

swagger-generate:
	@docker compose run --rm swagger \
		init \
		-g cmd/app/main.go \
		-o docs \
		--parseInternal \
		--parseDependency
