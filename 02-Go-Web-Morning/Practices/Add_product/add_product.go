package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required"`
	Code_value   string  `json:"code_value" binding:"required"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

func (p *Product) ValidateCodeValue() (message bool, err error) {
	for _, value := range products {
		if value.Code_value == p.Code_value {
			message = false
			err = errors.New("Invalid Code")
			return
		}
	}
	return true, nil
}

func (p *Product) ValidateExpirationDate() (bool, error) {
	date := strings.Split(p.Expiration, "/")
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])

	if (day > 0 && day < 32) && (month > 0 && month < 13) && (year > 2000) {
		return true, nil
	} else {
		return false, errors.New("Expiration date not valid")
	}

}

var products []Product
var lastId = 0

func main() {
	sv := gin.Default()
	sv.Get("/product/:id", GetProduct())

	sv.POST("/products", CreateProduct())

	// Run
	if err := sv.Run(":8080"); err != nil {
		panic(err)
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		param, _ = strconv.Atoi(param)
		product, err := SearchProduct(param)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err, "data": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    product,
		})
	}
}

func SearchProduct(param int) (Product, error) {
	for _, product := range products {
		if product.ID == param {
			return product, nil
		}
	}
	return Product{}, errors.New("ID not exist")
}

func CreateProduct() gin.HandlerFunc {
	type request struct {
		Name         string  `json:"name"`
		Quantity     int     `json:"quantity"`
		Code_value   string  `json:"code_value"`
		Is_published bool    `json:"is_published"`
		Expiration   string  `json:"expiration"`
		Price        float64 `json:"price"`
	}

	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"message": "invalid request", "data": nil,
			})
			return
		}

		product := Product{
			ID:           lastId + 1,
			Name:         req.Name,
			Quantity:     req.Quantity,
			Code_value:   req.Code_value,
			Is_published: req.Is_published,
			Expiration:   req.Expiration,
			Price:        req.Price,
		}

		if _, err := product.ValidateCodeValue(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err, "data": nil,
			})
			return
		}

		if _, err := product.ValidateExpirationDate(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err, "data": nil,
			})
			return
		}

		products = append(products, product)

		lastId++

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    product,
		})

	}
}
