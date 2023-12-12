package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practiva/web/internal/domain"
	"practiva/web/internal/product"
	"practiva/web/pkg"
	"strconv"
	"time"
)

type ProductController struct {
	productGroup *gin.RouterGroup
	service      product.ProductService
	repository   product.ProductRepository
}

func NewProductController(group *gin.RouterGroup) ProductController {
	products := pkg.FullfilDB("products.json")
	repository := product.NewProductRepository(products)
	service := product.NewProductService(repository)
	return ProductController{group, service, repository}
}
func (controller *ProductController) ProductRoutes() {
	controller.productGroup.GET("/ping", controller.Ping())
	controller.productGroup.GET("", controller.GetAllProducts())
	controller.productGroup.GET("/:id", controller.GetProductById())
	controller.productGroup.GET("/search", controller.GetAllWithPriceGreaterThan())
	controller.productGroup.POST("", controller.SaveProduct())
}
func (controller *ProductController) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	}
}
func isValidDate(date string) bool {
	_, err := time.Parse("02/01/2006", date)
	return err == nil
}
func (controller *ProductController) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := controller.service.GetAllProducts()
		ctx.JSON(http.StatusOK, data)
	}
}
func (controller *ProductController) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid id",
			})
			return
		}
		data := controller.service.GetById(int(id))
		if data.ID == 0 {
			ctx.AbortWithStatus(404)
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
func (controller *ProductController) GetAllWithPriceGreaterThan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid price",
			})
			return
		}
		data := controller.service.GetByPriceGreaterThan(price)
		ctx.JSON(http.StatusOK, data)
	}
}
func (controller *ProductController) SaveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newProduct domain.Product
		err := ctx.ShouldBindJSON(&newProduct)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid product",
			})
			return
		}
		if newProduct.Quantity < 0 {
			ctx.JSON(400, gin.H{
				"message": "invalid quantity",
			})
			return
		}
		if newProduct.Price < 0 {
			ctx.JSON(400, gin.H{
				"message": "invalid price",
			})
			return
		}
		if newProduct.Expiration == "" || !isValidDate(newProduct.Expiration) {
			ctx.JSON(400, gin.H{
				"message": "invalid expiration",
			})
			return
		}
		if newProduct.Name == "" {
			ctx.JSON(400, gin.H{
				"message": "invalid name",
			})
			return
		}
		if newProduct.CodeValue == "" || controller.service.ExistsAnyWithCodeValue(newProduct.CodeValue) {
			ctx.JSON(400, gin.H{
				"message": "invalid code value",
			})
			return
		}
		newProduct = controller.service.Save(newProduct)
		ctx.JSON(http.StatusCreated, newProduct)
	}
}
