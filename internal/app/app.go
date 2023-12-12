package app

import (
	"clean_architecture/internal/config"
	delivery "clean_architecture/internal/delivery/http"
	"clean_architecture/internal/repository"
	"clean_architecture/internal/server"
	"clean_architecture/internal/service"
	"clean_architecture/pkg/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	// Dependencies
	//TODO Пока закоментировано, потому что нет БД.
	//db, err := postgres.NewPostgresDB(cfg.Postgres)
	//if err != nil {
	//	logger.Error(err)
	//
	//	return
	//}

	// Services, Repos & API Handlers
	repos := repository.NewRepositories(nil)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := delivery.NewHandler(services)

	// HTTP Server
	srv := server.NewServer(cfg.Http, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
