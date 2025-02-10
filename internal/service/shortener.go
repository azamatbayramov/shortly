package service

import (
	"errors"
	"regexp"
	"shortly/config"

	"shortly/internal/appErrors"
	"shortly/internal/storage"
	"shortly/pkg/coder"
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

func (service ShortenerService) GetFullLink(shortLink string) (string, error) {
	id, err := service.coder.Decode(shortLink)

	if err != nil {
		return "", appErrors.ShortLinkIsNotValid
	}

	link, err := service.storage.GetLinkById(id)

	if err != nil {
		if errors.Is(err, appErrors.LinkNotFound) {
			return "", appErrors.LinkNotFound
		}

		return "", appErrors.StorageError
	}

	return link, nil
}

func (service ShortenerService) ShortenLink(link string) (string, error) {
	if len(link) > service.config.OriginalLinkMaxLength {
		return "", appErrors.OriginalLinkIsTooLong
	}

	var validLinkRegex = regexp.MustCompile(`^(http|https):\/\/[^\s/$.?#].[^\s]*$`)
	if !validLinkRegex.MatchString(link) {
		return "", appErrors.OriginalLinkIsNotValid
	}

	id, err := service.storage.GetIdByLinkOrAddNew(link)

	if err != nil {
		return "", appErrors.StorageError
	}

	shortLink, err := service.coder.Encode(id)

	if err != nil {
		return "", appErrors.EncodeError
	}

	return shortLink, nil
}
