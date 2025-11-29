package rest

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Util787/mws-content-registry/internal/config"
)

const (
	defaultReadTimeout       = 5 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultReadHeaderTimeout = 5 * time.Second
	defaultMaxHeaderBytes    = 1 << 20 // 1 MB
)

type Server struct {
	httpServer *http.Server
}

func NewRestServer(log *slog.Logger, config config.HTTPServerConfig, mwsusecase MWSTablesUsecase) Server { // add usecase in args
	handler := Handler{
		log:              log,
		MWSTablesUsecase: mwsusecase,
	}

	httpServer := &http.Server{
		Addr:              config.Host + ":" + strconv.Itoa(config.Port),
		Handler:           handler.InitRoutes(),
		MaxHeaderBytes:    defaultMaxHeaderBytes,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		WriteTimeout:      defaultWriteTimeout,
		ReadTimeout:       defaultReadTimeout,
	}

	return Server{
		httpServer: httpServer,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
