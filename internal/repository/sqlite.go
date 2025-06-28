package repository

import (
	"context"
	"mine-server-manager/internal/config"

	"github.com/jmoiron/sqlx"
)

type sqliteRepository struct {
	db *sqlx.DB
}

func NewSqliteRepository(ctx context.Context, cfg *config.Config) (*sqliteRepository, error) {
	db, err := sqlx.ConnectContext(ctx, cfg.SqlDriver, cfg.SqlConnString)
	if err != nil {
		return nil, err
	}

	return &sqliteRepository{
		db: db,
	}, nil
}

func (r *sqliteRepository) RetrieveHashedPwd(userEmail string) error {
	/// do impl later
	return nil
}
