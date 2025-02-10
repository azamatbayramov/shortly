package main

import (
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"

	"shortly/config"
	"shortly/internal/controller"
	"shortly/internal/service"
	"shortly/internal/storage"
	"shortly/pkg/coder"
)

func main() {
	appConfig, err := config.LoadConfig()

	if err != nil {
		slog.Error("failed to load config", "error", err)
		return
	}

	r := gin.Default()

	var appStorage storage.Storage

	if appConfig.StorageType == "in_memory" {
		appStorage, err = storage.NewMemoryStorage()
	} else if appConfig.StorageType == "postgresql" {
		appStorage, err = storage.NewPostgreSQLStorage(appConfig)
	}

	if err != nil {
		slog.Error("failed to create storage", "error", err)
		return
	}

	appCoder, err := coder.NewBaseCoder(appConfig.CoderAlphabet, appConfig.CoderLength)

	if err != nil {
		slog.Error("failed to create coder", "error", err)
		return
	}

	shortenerService := service.NewShortenerService(appStorage, appCoder, appConfig)
	shortenerController := controller.NewShortenerController(shortenerService)

	r.POST("/shorten", shortenerController.ShortenLink)
	r.GET("/:short_url", shortenerController.GetLink)

	err = r.Run(appConfig.AppHost + ":" + strconv.Itoa(appConfig.AppPort))
	if err != nil {
		slog.Error("failed to run server", "error", err)
		return
	}
}
