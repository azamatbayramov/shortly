package tests

import (
	"shortly/config"
	"shortly/internal/appErrors"
	"shortly/internal/service"
	"shortly/internal/storage"
	"shortly/pkg/coder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenerService_ShortenLink(t *testing.T) {
	store, _ := storage.NewMemoryStorage()
	cdr, _ := coder.NewBaseCoder("abcdefghijklmnopqrstuvwxyz", 6)
	cfg := &config.Config{OriginalLinkMaxLength: 100}
	shortener := service.NewShortenerService(store, cdr, cfg)

	link := "https://example.com"
	shortLink, err := shortener.ShortenLink(link)
	assert.NoError(t, err)
	assert.NotEmpty(t, shortLink)

	shortLink2, err := shortener.ShortenLink(link)
	assert.NoError(t, err)
	assert.Equal(t, shortLink, shortLink2)

	invalidLink := "invalid_link"
	_, err = shortener.ShortenLink(invalidLink)
	assert.ErrorIs(t, err, appErrors.OriginalLinkIsNotValid)
}

func TestShortenerService_GetFullLink(t *testing.T) {
	store, _ := storage.NewMemoryStorage()
	cdr, _ := coder.NewBaseCoder("abcdefghijklmnopqrstuvwxyz", 6)
	cfg := &config.Config{OriginalLinkMaxLength: 100}
	shortener := service.NewShortenerService(store, cdr, cfg)

	link := "https://example.com"
	shortLink, _ := shortener.ShortenLink(link)

	retrievedLink, err := shortener.GetFullLink(shortLink)
	assert.NoError(t, err)
	assert.Equal(t, link, retrievedLink)

	_, err = shortener.GetFullLink("invalid")
	assert.ErrorIs(t, err, appErrors.ShortLinkIsNotValid)

	_, err = shortener.GetFullLink("zzzzzz")
	assert.ErrorIs(t, err, appErrors.LinkNotFound)
}
