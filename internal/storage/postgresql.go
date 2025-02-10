package storage

import (
	"context"
	"fmt"
	"github.com/azamatbayramov/shortly/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
	"time"
)

type PostgreSQLStorage struct {
	pool *pgxpool.Pool
}

var _ Storage = (*PostgreSQLStorage)(nil)

func NewPostgreSQLStorage(config *config.Config) (*PostgreSQLStorage, error) {
	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		config.PsqlUser,
		config.PsqlPassword,
		config.PsqlHost,
		strconv.Itoa(config.PsqlPort),
		config.PsqlDatabase,
	)

	pool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		return nil, err
	}

	return &PostgreSQLStorage{pool: pool}, nil
}

func (stor PostgreSQLStorage) GetLinkById(id uint64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT link FROM links WHERE id = $1`

	var link string
	err := stor.pool.QueryRow(ctx, query, id).Scan(&link)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", ErrLinkNotFound
		}

		return "", err
	}

	return link, nil
}

func (stor PostgreSQLStorage) GetOrCreateLink(link string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getQuery := `SELECT id FROM links WHERE link = $1`

	var id uint64
	err := stor.pool.QueryRow(ctx, getQuery, link).Scan(&id)

	if err == nil {
		return id, nil
	}

	if err.Error() != "no rows in result set" {
		return 0, err
	}

	insertQuery := `
WITH inserted AS (
    INSERT INTO links (link) 
    VALUES ($1)
    ON CONFLICT (link) DO NOTHING
    RETURNING id
)
SELECT id FROM inserted 
UNION ALL
SELECT id FROM links WHERE link = $1
LIMIT 1;
`

	err = stor.pool.QueryRow(ctx, insertQuery, link).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
