package main

import (
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

	r := gin.Default()

	appStorage, err := storage.NewMemoryStorage()
	appCoder, err := coder.NewBaseCoder(appConfig.CoderAlphabet, appConfig.CoderLength)

	if err != nil {
		panic(err)
	}

	shortenerService := service.NewShortenerService(appStorage, appCoder, appConfig)
	shortenerController := controller.NewShortenerController(shortenerService)

	r.POST("/shorten", shortenerController.ShortenLink)
	r.GET("/:short_url", shortenerController.GetLink)

	err = r.Run(appConfig.AppHost + ":" + strconv.Itoa(appConfig.AppPort))
	if err != nil {
		return
	}
}
