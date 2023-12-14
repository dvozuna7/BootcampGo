package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"practiva/web/cmd/server/handler"
	"practiva/web/pkg"
	"practiva/web/pkg/impl"
)

func main() {
	router := gin.Default()
	productRepository := impl.NewSliceBasedRepository(pkg.FullfilDB("products.json"))
	productService := impl.NewDefaultService(&productRepository)
	productHandler := handler.NewProductHandler(&productService)
	group := router.Group("/products")
	{
		group.GET("", productHandler.GetAllProducts)
		group.GET("/:id", productHandler.GetProductById)
		group.GET("/search", productHandler.GetAllWithPriceGreaterThan)
		group.GET("/ping", productHandler.Ping)
	}
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
