package service

import (
	"errors"
	"github.com/azamatbayramov/shortly/config"
	"log/slog"
	"regexp"

	"github.com/azamatbayramov/shortly/internal/appErrors"
	"github.com/azamatbayramov/shortly/internal/storage"
	"github.com/azamatbayramov/shortly/pkg/coder"
)

type ShortenerService struct {
	storage storage.Storage
	coder   coder.Coder
	config  *config.Config
}

func NewShortenerService(storage storage.Storage, coder coder.Coder, config *config.Config) *ShortenerService {
	return &ShortenerService{
		storage: storage,
		coder:   coder,
		config:  config,
	}
}

func (srv ShortenerService) GetFullLink(shortLink string) (string, error) {
	id, err := srv.coder.Decode(shortLink)

	if err != nil {
		return "", appErrors.ShortLinkIsNotValid
	}

	link, err := srv.storage.GetLinkById(id)

	if err != nil {
		if errors.Is(err, appErrors.LinkNotFound) {
			return "", appErrors.LinkNotFound
		}

		slog.Error("failed to get link by id", "error", err)
		return "", appErrors.StorageError
	}

	return link, nil
}

func (srv ShortenerService) ShortenLink(link string) (string, error) {
	const linkRegexp = `^(http|https):\/\/[^\s/$.?#].[^\s]*$`
	if len(link) > srv.config.OriginalLinkMaxLength {
		return "", appErrors.OriginalLinkIsTooLong
	}

	var validLinkRegex = regexp.MustCompile(linkRegexp)
	if !validLinkRegex.MatchString(link) {
		return "", appErrors.OriginalLinkIsNotValid
	}

	id, err := srv.storage.GetOrCreateLink(link)

	if err != nil {
		slog.Error("failed to get id by link or add new", "error", err)
		return "", appErrors.StorageError
	}

	shortLink, err := srv.coder.Encode(id)

	if err != nil {
		slog.Error("failed to encode id", "error", err)
		return "", appErrors.EncodeError
	}

	return shortLink, nil
}
