package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear un router con gin
	router := gin.Default()

	// Captura la solicitud GET "/hello_world"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "PONG")
	})

	// Corremos nuestro servidro sobre el puerto 8080
	router.Run()
}
