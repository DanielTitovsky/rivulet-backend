package app_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	app_loger "github.com/DanielTitovsky/rivulet-backend.git/internal/app/loger"
	"go.uber.org/zap"
)

type HttpServer struct {
	mux    *http.ServeMux
	config HttpServerConfig
	log    *app_loger.Logger
}

func NewHttpServer(config HttpServerConfig, log *app_loger.Logger) *HttpServer {
	return &HttpServer{
		mux:    http.NewServeMux(),
		config: config,
		log:    log,
	}
}

func (h *HttpServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()

		h.log.Warn("start http server", zap.String("addr", h.config.Addr))

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown HTTP server...")

		shutdownContext, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTime,
		)
		defer cancel()

		if err := server.Shutdown(shutdownContext); err != nil {
			_ = server.Close()

			return fmt.Errorf("shotdown HTTP server: %w", err)
		}

		h.log.Warn("HTTP server stopped")

		return nil
	}

	return nil
}

func (h *HttpServer) RegisterRouters(apiRouters ...ApiVersinRouter) {
	for _, router := range apiRouters {
		prefix := "/api/" + string(router.ApiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}
