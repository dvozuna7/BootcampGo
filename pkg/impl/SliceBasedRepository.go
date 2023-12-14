package impl

import (
	"practiva/web/internal/domain"
)

type SliceBasedRepository struct {
	products *[]domain.Product
}

func NewSliceBasedRepository(products *[]domain.Product) SliceBasedRepository {
	return SliceBasedRepository{products}
}

func (r *SliceBasedRepository) FindAll() (products *[]domain.Product, err error) {
	return r.products, nil
}
func (r *SliceBasedRepository) FindById(id int) (products *domain.Product, err error) {
	for _, product := range *r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}
func (r *SliceBasedRepository) FindByPriceGreaterThan(price float64) (products *[]domain.Product, err error) {
	var result []domain.Product
	for i, prod := range *r.products {
		if prod.Price > price {
			result = append(result, (*r.products)[i])
		}
	}
	return &result, nil
}
func (r *SliceBasedRepository) Update(product *domain.Product) (products *domain.Product, err error) {
	for i, prod := range *r.products {
		if prod.ID == product.ID {
			(*r.products)[i] = *product
			return product, nil
		}
	}
	return nil, nil
}
