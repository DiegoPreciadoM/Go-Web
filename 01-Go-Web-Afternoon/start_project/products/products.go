package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func ReadText(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta daniado")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&products); err != nil {
		panic(err) // Retornar errores
	}

	fmt.Println(products)
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}

func ProductList(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func ProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	exist := false
	var product Product

	for _, current_product := range products {
		if current_product.Id == id {
			exist = true
			product = current_product
			break
		}
	}

	if exist == false {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product does not exist",
		})
	} else {
		c.JSON(http.StatusOK, product)
	}

}

func SearchProductByParam(c *gin.Context) {
	param, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Param",
		})
	}

	var products_response []Product

	for _, current_product := range products {
		if current_product.Price > param {
			products_response = append(products_response, current_product)
		}
	}

	if len(products_response) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No existen datos",
		})
	}

	c.JSON(http.StatusOK, products_response)

}

func main() {
	ReadText("products.json")

	server := gin.Default()
	server.GET("/ping", Pong)
	server.GET("/products", ProductList)
	server.GET("/products/:id", ProductById)
	server.GET("/products/search", SearchProductByParam)

	if err := http.ListenAndServe(":8080", server); err != nil {
		panic(err)
	}
}
