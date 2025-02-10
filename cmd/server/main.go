package main

import (
	"log/slog"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"

	"shortly/config"
	"shortly/internal/controller"
	"shortly/internal/service"
	"shortly/internal/storage"
	"shortly/pkg/coder"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		slog.Error("failed to load config", "error", err)
		return
	}

	r := gin.Default()

	var stor storage.Storage

	if cfg.StorageType == "in_memory" {
		stor, err = storage.NewMemoryStorage()
	} else if cfg.StorageType == "postgresql" {
		stor, err = storage.NewPostgreSQLStorage(cfg)
	}

	if err != nil {
		slog.Error("failed to create storage", "error", err)
		return
	}

	codr, err := coder.NewBaseCoder(cfg.CoderAlphabet, cfg.CoderLength)

	if err != nil {
		slog.Error("failed to create coder", "error", err)
		return
	}

	shortenerService := service.NewShortenerService(stor, codr, cfg)
	shortenerController := controller.NewShortenerController(shortenerService)

	r.POST("/shorten", shortenerController.ShortenLink)
	r.GET("/:short_url", shortenerController.GetLink)

	err = r.Run(net.JoinHostPort(cfg.AppHost, strconv.Itoa(cfg.AppPort)))
	if err != nil {
		slog.Error("failed to run server", "error", err)
		return
	}
}
