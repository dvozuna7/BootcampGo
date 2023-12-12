package product

import "practiva/web/internal/domain"

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{repo}
}

func (s *ProductService) GetAllProducts() []domain.Product {
	return s.repository.GetAllProducts()
}

func (s *ProductService) GetById(id int) domain.Product {
	return s.repository.GetById(id)
}

func (s *ProductService) GetByPriceGreaterThan(price float64) []domain.Product {
	products := s.repository.GetAllProducts()
	var result []domain.Product
	for i, prod := range products {
		if prod.Price >= price {
			result = append(result, products[i])
		}
	}
	return result
}
func (s *ProductService) ExistsAnyWithCodeValue(codeValue string) bool {
	return s.repository.existsAnyWithCodeValue(codeValue)
}
func (s *ProductService) Save(product domain.Product) domain.Product {
	return s.repository.save(product)
}
