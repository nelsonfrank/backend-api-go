package repository

import (
	"context"

	"github.com/nelsonfrank/backend-api-go/internal/db"
	"github.com/nelsonfrank/backend-api-go/internal/domain"
)

type userRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) domain.UserRepository {
	return &userRepository{q}
}

func (r *userRepository) GetByID(id int64) (domain.User, error) {
	u, err := r.q.GetUserByID(context.Background(), int32(id))
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: int64(u.ID), Name: u.Name, Email: u.Email}, nil
}

func (r *userRepository) Create(name, email string) (domain.User, error) {
	u, err := r.q.CreateUser(context.Background(), db.CreateUserParams{Name: name, Email: email})
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: int64(u.ID), Name: u.Name, Email: u.Email}, nil
}
