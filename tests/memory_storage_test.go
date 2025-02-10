package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/azamatbayramov/shortly/internal/storage"
)

func TestMemoryStorage_GetLinkById(t *testing.T) {
	store, _ := storage.NewMemoryStorage()

	_, err := store.GetLinkById(1)
	assert.ErrorIs(t, err, storage.ErrLinkNotFound)

	link := "https://example.com"
	id, _ := store.GetOrCreateLink(link)
	retrievedLink, err := store.GetLinkById(id)
	assert.NoError(t, err)
	assert.Equal(t, link, retrievedLink)
}

func TestMemoryStorage_GetIdByLinkOrAddNew(t *testing.T) {
	store, _ := storage.NewMemoryStorage()

	link := "https://example.com"
	id1, err := store.GetOrCreateLink(link)
	assert.NoError(t, err)

	id2, err := store.GetOrCreateLink(link)
	assert.NoError(t, err)
	assert.Equal(t, id1, id2)

	link2 := "https://example.org"
	id3, err := store.GetOrCreateLink(link2)
	assert.NoError(t, err)
	assert.NotEqual(t, id1, id3)
}
