package server

import (
	"context"
	"go-rest/internal/config"
	"go-rest/internal/handler"
	"go-rest/internal/models/validators"
	"go-rest/internal/repository"
	"go-rest/internal/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewServer(cfg *config.Config, db *sqlx.DB, logger *logrus.Logger) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr: s.cfg.Server.Port,
	}

	go func() {
		s.logger.Infof("Server is listening on port %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error during start server: %s", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return s.echo.Server.Shutdown(ctx)
}

func (s *Server) MapHandlers(e *echo.Echo) error {
	urlRepo := repository.NewUrlRepository(s.db)
	urlVal := validators.NewUrlValidator()
	urlUc := usecase.NewUrlUsecase(s.logger, s.cfg, urlRepo)
	urlHs := handler.NewUrlHandler(s.cfg, s.logger, urlUc, urlVal)
	v1 := e.Group("/api")
	urlGroup := v1.Group("")
	handler.MapRoutes(urlGroup, urlHs)
	return nil
}
