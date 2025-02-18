package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/azamatbayramov/shortly/config"
	"github.com/azamatbayramov/shortly/internal/service"
	"github.com/azamatbayramov/shortly/internal/storage"
	"github.com/azamatbayramov/shortly/pkg/coder"
	"github.com/azamatbayramov/shortly/pkg/link/validator"
)

func TestShortenerService_ShortenLink(t *testing.T) {
	store, _ := storage.NewMemoryStorage()
	cdr, _ := coder.NewBaseCoder("abcdefghijklmnopqrstuvwxyz", 6)
	vldr := validator.NewNetUrlValidator()
	cfg := &config.Config{OriginalLinkMaxLength: 100}
	shortener := service.NewShortenerService(store, cdr, vldr, cfg)

	link := "https://example.com"
	shortLink, err := shortener.ShortenLink(link)
	assert.NoError(t, err)
	assert.NotEmpty(t, shortLink)

	shortLink2, err := shortener.ShortenLink(link)
	assert.NoError(t, err)
	assert.Equal(t, shortLink, shortLink2)

	invalidLink := "invalid_link"
	_, err = shortener.ShortenLink(invalidLink)
	assert.ErrorIs(t, err, service.ErrOriginalLinkIsNotValid)
}

func TestShortenerService_GetFullLink(t *testing.T) {
	store, _ := storage.NewMemoryStorage()
	cdr, _ := coder.NewBaseCoder("abcdefghijklmnopqrstuvwxyz", 6)
	vldr := validator.NewNetUrlValidator()
	cfg := &config.Config{OriginalLinkMaxLength: 100}
	shortener := service.NewShortenerService(store, cdr, vldr, cfg)

	link := "https://example.com"
	shortLink, _ := shortener.ShortenLink(link)

	retrievedLink, err := shortener.GetFullLink(shortLink)
	assert.NoError(t, err)
	assert.Equal(t, link, retrievedLink)

	_, err = shortener.GetFullLink("invalid")
	assert.ErrorIs(t, err, service.ErrShortLinkIsNotValid)

	_, err = shortener.GetFullLink("zzzzzz")
	assert.ErrorIs(t, err, service.ErrLinkNotFound)
}
