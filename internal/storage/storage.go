package storage

type Storage interface {
	GetLinkById(id uint64) (string, error)
	GetOrCreateLink(link string) (uint64, error)
}
