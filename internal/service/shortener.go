package service

import (
	"errors"
	"log/slog"

	"github.com/azamatbayramov/shortly/config"
	"github.com/azamatbayramov/shortly/internal/storage"
	"github.com/azamatbayramov/shortly/pkg/coder"
	"github.com/azamatbayramov/shortly/pkg/link/validator"
)

type ShortenerService struct {
	storage       storage.Storage
	coder         coder.Coder
	linkValidator validator.Validator
	config        *config.Config
}

func NewShortenerService(storage storage.Storage, coder coder.Coder, linkValidator validator.Validator, config *config.Config) *ShortenerService {
	return &ShortenerService{
		storage:       storage,
		coder:         coder,
		linkValidator: linkValidator,
		config:        config,
	}
}

func (srv ShortenerService) GetFullLink(shortLink string) (string, error) {
	id, err := srv.coder.Decode(shortLink)

	if err != nil {
		return "", ErrShortLinkIsNotValid
	}

	link, err := srv.storage.GetLinkById(id)

	if err != nil {
		if errors.Is(err, storage.ErrLinkNotFound) {
			return "", ErrLinkNotFound
		}

		slog.Error("failed to get link by id", "error", err)
		return "", ErrStorageError
	}

	return link, nil
}

func (srv ShortenerService) ShortenLink(link string) (string, error) {
	if len(link) > srv.config.OriginalLinkMaxLength {
		return "", ErrOriginalLinkIsTooLong
	}

	if !srv.linkValidator.Validate(link) {
		return "", ErrOriginalLinkIsNotValid
	}

	id, err := srv.storage.GetOrCreateLink(link)

	if err != nil {
		slog.Error("failed to get id by link or add new", "error", err)
		return "", ErrStorageError
	}

	shortLink, err := srv.coder.Encode(id)

	if err != nil {
		slog.Error("failed to encode id", "error", err)
		return "", ErrEncodeError
	}

	return shortLink, nil
}
