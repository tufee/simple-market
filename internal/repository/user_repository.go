package repository

import (
	"simple.market/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id int) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int) error
}
