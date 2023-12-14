package impl

import (
	"practiva/web/internal/domain"
	"practiva/web/internal/product"
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
func (r *SliceBasedRepository) FindById(id int) (productById *domain.Product, err error) {
	for _, currentProduct := range *r.products {
		if currentProduct.ID == id {
			return &currentProduct, nil
		}
	}
	return nil, product.ProductNotFound
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
func (r *SliceBasedRepository) DeleteById(id int) error {
	searchedProduct, err := r.FindById(id)
	if err != nil {
		return err
	}
	if searchedProduct == nil {
		return product.ProductNotFound
	}
	for i, prod := range *r.products {
		if prod.ID == id {
			*r.products = append((*r.products)[:i], (*r.products)[i+1:]...)
			break
		}
	}
	return nil
}
