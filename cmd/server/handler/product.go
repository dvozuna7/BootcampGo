package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"practiva/web/internal/product"
	"strconv"
	"time"
)

type ProductHandler struct {
	service product.Service
}

func NewProductHandler(s product.Service) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (h *ProductHandler) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pong")
}
func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.service.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected error",
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}
	productById, err := h.service.FindById(ctx, int(id))
	if err != nil {
		switch err {
		case product.ProductNotFound:
			ctx.JSON(http.StatusNotFound, err)
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "unexpected error",
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, productById)
}
func (h *ProductHandler) GetAllWithPriceGreaterThan(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid price",
		})
		return
	}
	products, err := h.service.FindByPriceGreaterThan(ctx, price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected error",
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
func (h *ProductHandler) DeleteProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}
	fmt.Println("TODO: ADD DELETION: ", id)
	//TODO: Add deletion to service
}
func isValidDate(date string) bool {
	_, err := time.Parse("02/01/2006", date)
	return err == nil
}
