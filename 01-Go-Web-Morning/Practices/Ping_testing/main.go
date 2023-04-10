package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.Data(200, "application/json; charset=utf-8", []byte("PONG"))
	})

	router.Run(":8080")

}
