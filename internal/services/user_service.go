package services

import "github.com/nelsonfrank/backend-api-go/internal/domain"

type UserService struct {
	Repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{r}
}

func (s *UserService) RegisterUser(name, email string) (domain.User, error) {
	return s.Repo.Create(name, email)
}
