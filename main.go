package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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
	return file, err
}
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
func populateProducts() error {
	file, err := openFile("products.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer closeFile(file)
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
func getNewId() int {
	maxId := 0
	for _, product := range products {
		if product.ID > maxId {
			maxId = product.ID
		}
	}
	return maxId + 1
}
func existsAnyWithCodeValue(codeValue string) bool {
	for _, product := range products {
		if product.CodeValue == codeValue {
			return true
		}
	}
	return false
}
func isValidDate(date string) bool {
	_, err := time.Parse("02/01/2006", date)
	return err == nil
}
func SaveProduct(ctx *gin.Context) {
	var product Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid product",
		})
		return
	}
	product.ID = getNewId()
	if product.Quantity < 0 {
		ctx.JSON(400, gin.H{
			"message": "invalid quantity",
		})
		return
	}
	if product.Price < 0 {
		ctx.JSON(400, gin.H{
			"message": "invalid price",
		})
		return
	}
	if product.Expiration == "" || isValidDate(product.Expiration) {
		ctx.JSON(400, gin.H{
			"message": "invalid expiration",
		})
		return
	}
	if product.Name == "" {
		ctx.JSON(400, gin.H{
			"message": "invalid name",
		})
		return
	}
	if product.CodeValue == "" || existsAnyWithCodeValue(product.CodeValue) {
		ctx.JSON(400, gin.H{
			"message": "invalid code value",
		})
		return
	}
	products = append(products, product)
	ctx.JSON(http.StatusCreated, gin.H{
		"product": product,
	})

}

func main() {
	engine := gin.Default()
	err := populateProducts()
	if err != nil {
		panic(err)
	}
	engine.GET("/ping", Ping)
	engine.GET("/products", GetProducts)
	engine.GET("/products/:id", GetProductById)
	engine.GET("/products/search", GetAllWithPriceGreaterThan)
	engine.POST("/products", SaveProduct)
	err = engine.Run()
	if err != nil {
		fmt.Println(err)
	}
}
