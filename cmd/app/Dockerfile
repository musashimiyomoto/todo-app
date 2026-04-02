# Build stage
FROM golang:1.25-bookworm AS builder

WORKDIR /todo-app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app/cmd/app/exe /todo-app/cmd/app 

# Final stage
FROM alpine:3.23

WORKDIR /todo-app

COPY --from=builder /todo-app/cmd/app/exe /todo-app
COPY --from=builder /todo-app/public /todo-app/public

ENV PROJECT_ROOT=/todo-app

CMD ["/todo-app/exe"]
