package repository

import "context"

type SQLRepository interface {
	CreateUser(ctx context.Context, userEmail, userName, pwd string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type sqlRepository struct {
	*Queries
}

func NewSQLRepository(db DBTX) SQLRepository {
	return &sqlRepository{
		Queries: New(db),
	}
}

func (r *sqlRepository) CreateUser(ctx context.Context, userEmail, userName, pwd string) (User, error) {
	return r.Queries.CreateUser(ctx, CreateUserParams{
		Username: userName,
		Email:    userEmail,
		Password: pwd,
	})
}

func (r *sqlRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return r.Queries.GetUserByEmail(ctx, email)
}
