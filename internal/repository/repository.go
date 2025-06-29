package repository

type Repository struct {
	SqlRepo         SQLRepository
	NoSqlRepository NoSqlRepository
	CacheRepository CacheRepository
}
