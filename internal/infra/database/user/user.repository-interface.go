package database

import "github.com/eltoncasacio/api-go/internal/entity"

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
