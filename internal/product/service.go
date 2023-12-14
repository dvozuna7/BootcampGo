package product

import (
	"github.com/gin-gonic/gin"
	"practiva/web/internal/domain"
)

type Service interface {
	FindAll(ctx *gin.Context) (*[]domain.Product, error)
	FindById(ctx *gin.Context, id int) (*domain.Product, error)
	FindByPriceGreaterThan(ctx *gin.Context, price float64) (*[]domain.Product, error)
	Update(ctx *gin.Context, product *domain.Product) (*domain.Product, error)
}
