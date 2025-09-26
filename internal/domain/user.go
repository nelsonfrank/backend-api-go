package domain

type User struct {
	ID    int64
	Name  string
	Email string
}

type UserRepository interface {
	GetByID(id int64) (User, error)
	Create(name, email string) (User, error)
}
