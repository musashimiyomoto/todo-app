# Todo App API

REST API for managing tasks, users, and statistics.

## Requirements

- Go 1.25+
- Docker, Docker Compose
- Make

## Setup

```bash
cp .env.example .env
```

Edit `.env` if needed.

## Run with Docker

```bash
# Start all services (app + postgres)
make up

# Apply migrations
make migrate-up

# Stop
make down
```

## Run locally

A running PostgreSQL instance is required (can be started via Docker):

```bash
# Start only postgres + socat (proxies port to localhost:5432)
docker compose up -d postgres socat

# Apply migrations
make migrate-up

# Run the application
make run
```

API is available at `http://127.0.0.1:5050`.  
Swagger UI: `http://127.0.0.1:5050/swagger/`.

## Migrations

```bash
make migrate-up                    # apply all migrations
make migrate-down                  # rollback last migration
make migrate-create seq=add_index  # create a new migration
```

## Other

```bash
make swagger-generate  # regenerate swagger docs
make cleanup           # clean postgres data and logs
```
