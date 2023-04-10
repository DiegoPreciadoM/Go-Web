package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	product "github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Morning/Practices/Full_practice/Products"
	"github.com/gin-gonic/gin"
)

var currentId int
var sv product.SaveProduct

func main() {
	if err := sv.GetProducts("Products/products.json"); err != nil {
		panic(err)
	}

	server := gin.Default()
	server.POST("/products", CreateNewProduct())
	server.GET("/products/:id", GetProductById())

	server.Run(":8080")
}

func GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id_param, _ := strconv.Atoi(ctx.Param("id"))

		data_product, err := SearchProduct(id_param)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"data": data_product,
		})
		return
	}
}

func SearchProduct(id int) (product.Product, error) {
	data := sv.GetInformationProducts()
	for _, value := range data {
		if value.Id == id {
			return value, nil
		}
	}
	return product.Product{}, errors.New("Product does not exist")
}

func CreateNewProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var new_product product.Product

		if err := ctx.ShouldBindJSON(&new_product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Verify body information",
			})
			return
		}

		if !ValidateCodeValue(new_product) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Code Value already exist",
			})
			return
		}

		if !ValidateExpirationDate(new_product) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Expiration date is incorrect",
			})
			return
		}

		currentId = len(sv.GetInformationProducts())
		new_product.Id = currentId + 1
		sv.SaveNewProduct(new_product)

		ctx.JSON(http.StatusCreated, gin.H{
			"data": new_product,
		})
		return

	}
}

func ValidateCodeValue(new_product product.Product) bool {
	data := sv.GetInformationProducts()
	for _, value := range data {
		if value.Code_value == new_product.Code_value {
			return false
		}
	}
	return true
}

func ValidateExpirationDate(new_product product.Product) bool {
	splited_date := strings.Split(new_product.Expiration, "/")
	day, _ := strconv.Atoi(splited_date[0])
	month, _ := strconv.Atoi(splited_date[1])
	year, _ := strconv.Atoi(splited_date[2])

	if day < 1 || day > 31 {
		return false
	}

	if month < 1 || month > 12 {
		return false
	}

	if year < 0 || year > 9999 {
		return false
	}

	return true
}
