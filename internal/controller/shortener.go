package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/azamatbayramov/shortly/internal/appErrors"
	"github.com/azamatbayramov/shortly/internal/models"
	"github.com/azamatbayramov/shortly/internal/service"
)

type ShortenerController struct {
	service *service.ShortenerService
}

func NewShortenerController(service *service.ShortenerService) *ShortenerController {
	return &ShortenerController{service: service}
}

func (ctrl ShortenerController) ShortenLink(c *gin.Context) {
	var link models.FullLink

	if err := c.ShouldBind(&link); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	shortLink, err := ctrl.service.ShortenLink(link.FullLink)

	if err != nil {
		if errors.Is(err, appErrors.OriginalLinkIsNotValid) || errors.Is(err, appErrors.OriginalLinkIsTooLong) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ShortLink{ShortLink: shortLink})
}

func (ctrl ShortenerController) GetLink(c *gin.Context) {
	shortLink := c.Param("short_url")

	fullLink, err := ctrl.service.GetFullLink(shortLink)

	if err != nil {
		if errors.Is(err, appErrors.ShortLinkIsNotValid) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, appErrors.LinkNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, fullLink)
}
