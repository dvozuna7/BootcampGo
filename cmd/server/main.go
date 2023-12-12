package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"practiva/web/cmd/server/controller"
)

func main() {
	server := gin.Default()
	group := server.Group("/products")
	router := controller.NewProductController(group)
	router.ProductRoutes()
	err := server.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
