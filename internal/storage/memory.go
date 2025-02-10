package storage

import (
	"shortly/internal/appErrors"
	"sync"
)

type MemoryStorage struct {
	idToLink map[uint64]string
	linkToId map[string]uint64
	count    uint64

	mu sync.RWMutex
}

var _ Storage = (*MemoryStorage)(nil)

func NewMemoryStorage() (*MemoryStorage, error) {
	return &MemoryStorage{
		idToLink: make(map[uint64]string),
		linkToId: make(map[string]uint64),
	}, nil
}

func (storage *MemoryStorage) GetLinkById(id uint64) (string, error) {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	link, exists := storage.idToLink[id]
	if !exists {
		return "", appErrors.LinkNotFound
	}

	return link, nil
}

func (storage *MemoryStorage) GetIdByLinkOrAddNew(link string) (uint64, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	id, exists := storage.linkToId[link]
	if exists {
		return id, nil
	}

	id = storage.count

	storage.count++
	storage.idToLink[id] = link
	storage.linkToId[link] = id

	return id, nil
}
