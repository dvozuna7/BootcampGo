package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int64   `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func openFile(fileName string) (file *os.File, err error) {
	file, err = os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func populateProducts() error {
	file, err := openFile("products.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	productsBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(productsBytes, &products)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"products": products,
	})
}
func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}
	for _, product := range products {
		if product.ID == id {
			c.JSON(200, gin.H{
				"product": product,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "product not found",
	})
}
func GetAllWithPriceGreaterThan(c *gin.Context) {
	price, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid price",
		})
		return
	}
	var filteredProducts []Product
	for _, product := range products {
		if product.Price > price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	c.JSON(200, gin.H{
		"products": filteredProducts,
	})
}

func main() {
	// gin default
	engine := gin.Default()

	// gin router
	err := populateProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	engine.GET("/ping", Ping)
	engine.GET("/products", GetProducts)
	engine.GET("/products/:id", GetProductById)
	engine.GET("/products/search", GetAllWithPriceGreaterThan)

	err = engine.Run()
	if err != nil {
		fmt.Println(err)
	}
}
