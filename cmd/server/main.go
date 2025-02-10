package main

import (
	"log/slog"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/azamatbayramov/shortly/config"
	"github.com/azamatbayramov/shortly/internal/controller"
	"github.com/azamatbayramov/shortly/internal/service"
	"github.com/azamatbayramov/shortly/internal/storage"
	"github.com/azamatbayramov/shortly/pkg/coder"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg, err := config.LoadConfig()

	if err != nil {
		slog.Error("failed to load config", "error", err)
		return
	}

	r := gin.Default()

	r.LoadHTMLFiles("html/index.html")

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
	r.GET("/", shortenerController.GetMainPage)

	err = r.Run(net.JoinHostPort(cfg.AppHost, strconv.Itoa(cfg.AppPort)))
	if err != nil {
		slog.Error("failed to run server", "error", err)
		return
	}
}
