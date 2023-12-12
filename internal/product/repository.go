package product

import (
	"practiva/web/internal/domain"
	"practiva/web/pkg"
)

type ProductRepository struct {
	productsDB []domain.Product
}

func NewProductRepository(products []domain.Product) ProductRepository {
	return ProductRepository{products}
}

func (r ProductRepository) GetAllProducts() []domain.Product {
	return r.productsDB
}

func (r ProductRepository) GetById(id int) domain.Product {
	for _, product := range r.GetAllProducts() {
		if product.ID == id {
			return product
		}
	}
	return domain.Product{}
}
func (r ProductRepository) ExistsById(id int) bool {
	return pkg.Any(r.productsDB, func(product domain.Product) bool {
		return product.ID == id
	})
}

func getNewId(r ProductRepository) int {
	maxId := 0
	for _, product := range r.productsDB {
		if product.ID > maxId {
			maxId = product.ID
		}
	}
	return maxId + 1
}
func (r *ProductRepository) save(product domain.Product) domain.Product {
	if r.ExistsById(product.ID) {
		for i, p := range r.productsDB {
			if p.ID == product.ID {
				r.productsDB[i] = product
				return product
			}
		}
	}
	product.ID = getNewId(*r)
	r.productsDB = append(r.productsDB, product)
	return product
}
func (r ProductRepository) existsAnyWithCodeValue(codeValue string) bool {
	return pkg.Any(r.productsDB, func(product domain.Product) bool {
		return product.CodeValue == codeValue
	})
}
