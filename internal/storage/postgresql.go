package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"shortly/config"
	"shortly/internal/appErrors"
	"strconv"
	"time"
)

type PostgreSQLStorage struct {
	pool *pgxpool.Pool
}

var _ Storage = (*PostgreSQLStorage)(nil)

func NewPostgreSQLStorage(config *config.Config) (*PostgreSQLStorage, error) {
	var pool *pgxpool.Pool
	var err error

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		strconv.Itoa(config.PostgresPort),
		config.PostgresDatabase,
	)

	pool, err = pgxpool.New(context.Background(), dbSource)

	if err != nil {
		return nil, err
	}

	return &PostgreSQLStorage{pool: pool}, nil
}

func (storage PostgreSQLStorage) GetLinkById(id uint64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var link string

	query := `SELECT link FROM links WHERE id = $1`

	err := storage.pool.QueryRow(ctx, query, id).Scan(&link)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", appErrors.LinkNotFound
		}

		return "", err
	}

	return link, nil
}

func (storage PostgreSQLStorage) GetIdByLinkOrAddNew(link string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id uint64

	query := `SELECT id FROM links WHERE link = $1
UNION ALL
WITH inserted AS (
    INSERT INTO links (link) 
    VALUES ($1)
    ON CONFLICT (link) DO NOTHING
    RETURNING id
)
SELECT id FROM inserted
LIMIT 1;
`

	err := storage.pool.QueryRow(ctx, query, link).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
