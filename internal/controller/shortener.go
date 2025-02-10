package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shortly/internal/models"
	"shortly/internal/service"
)

type ShortenerController struct {
	service *service.ShortenerService
}

func NewShortenerController(service *service.ShortenerService) *ShortenerController {
	return &ShortenerController{service: service}
}

func (ctrl ShortenerController) ShortenLink(c *gin.Context) {
	var link models.ShortenLink

	if err := c.ShouldBind(&link); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	shortLink, err := ctrl.service.ShortenLink(link.FullLink)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ShortenedLink{ShortLink: shortLink})
}

func (ctrl ShortenerController) GetLink(c *gin.Context) {
	shortLink := c.Param("short_url")

	fullLink, err := ctrl.service.GetFullLink(shortLink)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, fullLink)
}
