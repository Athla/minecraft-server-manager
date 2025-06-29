package repository

type NoSqlRepository interface {
	Fetch(id string) (any, error)
	Insert(item any) error
}
