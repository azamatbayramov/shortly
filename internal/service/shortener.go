package service

import (
	"shortly/internal/storage"
	"shortly/pkg/coder"
)

type ShortenerService struct {
	storage storage.Storage
	coder   coder.Coder
}

func NewShortenerService(storage storage.Storage, coder coder.Coder) *ShortenerService {
	return &ShortenerService{
		storage: storage,
		coder:   coder,
	}
}

func (service ShortenerService) GetFullLink(shortLink string) (string, error) {
	id, err := service.coder.Decode(shortLink)

	if err != nil {
		return "", err
	}

	return service.storage.GetLinkById(id)
}

func (service ShortenerService) ShortenLink(link string) (string, error) {
	id, err := service.storage.GetIdByLinkOrAddNew(link)

	if err != nil {
		return "", err
	}

	return service.coder.Encode(id)
}
