package product

import (
	"practiva/web/internal/domain"
)

type Repository interface {
	FindAll() (*[]domain.Product, error)
	FindById(id int) (*domain.Product, error)
	FindByPriceGreaterThan(price float64) (*[]domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
}
