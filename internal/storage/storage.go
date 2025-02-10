package storage

type Storage interface {
	GetLinkById(id uint64) (string, error)
	GetIdByLinkOrAddNew(link string) (uint64, error)
}
