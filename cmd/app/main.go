package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	core_postgres_pool "github.com/musashimiyomoto/todo-app/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/musashimiyomoto/todo-app/internal/core/transport/http/middleware"
	core_http_server "github.com/musashimiyomoto/todo-app/internal/core/transport/http/server"
	users_postgres_repository "github.com/musashimiyomoto/todo-app/internal/features/users/repository/postgres"
	users_service "github.com/musashimiyomoto/todo-app/internal/features/users/service"
	users_transport_http "github.com/musashimiyomoto/todo-app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func initHTTPServer(logger *core_logger.Logger, pool *core_postgres_pool.ConnectionPool) *core_http_server.HTTPServer {
	logger.Debug("Initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("Initializing HTTP server...")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	return httpServer
}

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger")
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initializing database connection pool...")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("Failed to init database connection pool", zap.Error(err))
	}
	defer pool.Close()

	httpServer := initHTTPServer(logger, pool)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
