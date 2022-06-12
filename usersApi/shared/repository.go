package shared

import "usersApi/domain"

type Repository interface {
	Migrate() error
	Create(user domain.User) (*domain.User, error)
	All() ([]domain.User, error)
	GetById(id int64) (*domain.User, error)
	GetByName(name string) (*domain.User, error)
	Update(id int64, updated domain.User) (*domain.User, error)
	Delete(id int64) error
	Close() error
}
