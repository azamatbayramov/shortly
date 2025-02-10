package storage

import (
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

func (stor *MemoryStorage) GetLinkById(id uint64) (string, error) {
	stor.mu.RLock()
	defer stor.mu.RUnlock()

	link, exists := stor.idToLink[id]
	if !exists {
		return "", ErrLinkNotFound
	}

	return link, nil
}

func (stor *MemoryStorage) GetOrCreateLink(link string) (uint64, error) {
	stor.mu.Lock()
	defer stor.mu.Unlock()

	id, exists := stor.linkToId[link]
	if exists {
		return id, nil
	}

	id = stor.count

	stor.count++
	stor.idToLink[id] = link
	stor.linkToId[link] = id

	return id, nil
}
