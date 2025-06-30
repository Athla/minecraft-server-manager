package repository

import (
	"mine-server-manager/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	SqlRepo         SQLRepository
	NoSqlRepository NoSqlRepository
	CacheRepository CacheRepository
}

func NewRepository(cfg *config.Config) *Repository {
	dbtx, err := sqlx.Open("sqlite3", cfg.SqlConnString)
	if err != nil {
		panic(err)
	}
	return &Repository{
		SqlRepo:         NewSQLRepository(dbtx),
		CacheRepository: NewInMemoryCache(),
	}
}
