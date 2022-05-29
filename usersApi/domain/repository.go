package domain

type Repository interface {
	Migrate() error
	Create(user User) (*User, error)
	All() ([]User, error)
	GetByName(name string) (*User, error)
	Update(id int64, updated User) (*User, error)
	Delete(id int64) error
}
