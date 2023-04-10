package main

import (
	"log"
	"net/http"
	"strconv"

	products "github.com/DiegoPreciadoM/Go-Web/tree/main/01-Go-Web-Afternoon/Practices/Full_practice/Products"
	"github.com/gin-gonic/gin"
)

var sv products.SaveProduct

func main() {
	// Get information from products.json file
	var err error

	err = sv.GetProducts("Products/products.json")
	if err != nil {
		panic(err)
	}

	// Create Server
	server := gin.Default()

	server.GET("/ping", Ping())

	product := server.Group("/products")
	product.GET("", GetAllProducts())
	product.GET("/:id", GetProductById())
	product.GET("/search", GetProductByParam())

	server.Run(":8080")
}

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})
		return
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"data": sv.GetInformationProducts(),
		})
		return
	}
}

func GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		data := sv.GetInformationProducts()

		for _, value := range data {
			if value.Id == idParam {
				ctx.JSON(http.StatusAccepted, gin.H{
					"product": value,
				})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product does not exist",
		})
		return
	}
}

func GetProductByParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			log.Fatal(err)
		}
		data := sv.GetInformationProducts()
		var response_data []products.Product

		for _, value := range data {
			if value.Price > param {
				response_data = append(response_data, value)
			}
		}

		if len(response_data) > 0 {
			ctx.JSON(http.StatusAccepted, gin.H{
				"data": response_data,
			})
			return
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product does not exist",
		})
		return
	}
}
