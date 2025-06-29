package repository

type CacheRepository interface {
	Add(key string, value any) error
	Get(key string) (any, error)
}
