package database

import "github.com/eltoncasacio/api-go/internal/entity"

type ProductRepositoryInterface interface {
	Create(product *entity.Product) error
	Paginate(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
