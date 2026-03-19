package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_loger "github.com/musashimiyomoto/todo-app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	log    *core_loger.Logger
}

func NewHTTPServer(config Config, log *core_loger.Logger) *HTTPServer {
	return &HTTPServer{
		mux:    &http.ServeMux{},
		config: config,
		log:    log,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api" + string(router.apiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router.mux),
		)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		h.log.Warn("Start HTTP Server", zap.String("Addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and serve HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("Shutdown HTTP server...")

		shutDownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutDownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("Shutdown HTTP server: %w", err)
		}

		h.log.Warn("HTTP server stopped")
	}

	return nil
}
