package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fevse/effm/internal/app"
	"github.com/fevse/effm/internal/config"
	"github.com/fevse/effm/internal/logger"
	"github.com/fevse/effm/internal/server"
	"github.com/fevse/effm/internal/storage"
)

func main() {
	config := config.LoadConfig()
	logger := logger.NewLogger(config)
	storage := storage.NewStorage(config, logger)
	if err := storage.Connect(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer func() {
		err := storage.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	if err := storage.Migrate(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := app.NewEffmApp(storage, logger)

	server := server.NewServer(app, config.ServHost, config.ServPort)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Error(err.Error())
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := server.Start(ctx)
		if err != nil {
			logger.Error(err.Error())
			cancel()
		}
	}()
	<-ctx.Done()
	wg.Wait()
}
