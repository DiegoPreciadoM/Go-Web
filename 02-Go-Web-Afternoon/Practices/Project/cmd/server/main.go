package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/cmd/handler"
	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/internal/domain"
	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {
	// if err := sv.GetProducts("Products/products.json"); err != nil {
	// 	panic(err)
	// }

	// err := godotenv.Loaad()
	// if err != nil {
	// 	log.Fatal("Error loading token")
	// }

	var productsList = []domain.Product{}
	loadProducts("products.json", &productsList)

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	server := gin.Default()

	server.GET("ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "PONG") })
	products := server.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET("/:id", productHandler.GetById())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.DELETE("/:id", productHandler.Delete())
		products.PATCH("/:id", productHandler.Patch())
		products.PUT("/:id", productHandler.Put())
	}

	server.Run(":8080")
}

func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}
