package impl

import (
	"github.com/gin-gonic/gin"
	"practiva/web/internal/domain"
	"practiva/web/internal/product"
)

type DefaultService struct {
	Repository product.Repository
}

func NewDefaultService(repository product.Repository) DefaultService {
	return DefaultService{repository}
}

func (s *DefaultService) FindAll(ctx *gin.Context) (*[]domain.Product, error) {
	return s.Repository.FindAll()
}
func (s *DefaultService) FindById(ctx *gin.Context, id int) (*domain.Product, error) {
	productById, err := s.Repository.FindById(id)
	if err != nil {
		return nil, err
	}
	if productById == nil {
		return nil, nil
	}
	return productById, nil
}
func (s *DefaultService) FindByPriceGreaterThan(ctx *gin.Context, price float64) (*[]domain.Product, error) {
	return s.Repository.FindByPriceGreaterThan(price)
}
func (s *DefaultService) Update(ctx *gin.Context, product *domain.Product) (*domain.Product, error) {
	return s.Repository.Update(product)
}
