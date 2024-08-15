package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/martinyonatann/acte/config"
	"github.com/martinyonatann/acte/internal/middleware"
	"github.com/martinyonatann/acte/pkg/datasources"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Server struct {
	db   *gorm.DB
	echo *echo.Echo
	cfg  config.Config
}

func NewServer(ctx context.Context, cfg config.Config) *Server {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	db, err := datasources.NewMySQLDB(ctx, cfg.Database)
	if err != nil {
		panic(err)
	}

	return &Server{
		db:   db,
		echo: middleware.NewEcho(cfg),
		cfg:  cfg,
	}
}

func (s *Server) Run() error {
	if err := s.mapHandlers(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	// Shutdown gracefully.
	go func() {
		<-quit

		log.Info("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		lo.Must(s.db.DB()).Close()
		s.echo.Shutdown(ctx)
	}()

	return s.echo.Start(fmt.Sprintf(":%s", s.cfg.Server.Port))
}
