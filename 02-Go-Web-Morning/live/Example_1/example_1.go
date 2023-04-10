package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Env
	//...

	// Instances
	// ...

	// Server
	sv := gin.Default()
	sv.Get("/ping", func(c *gin.Context) {
		//c.JSON(200, gin.H{"data": "pong"})
		c.string(200, "pong")
	})

	sv.POST("/movies", CreateMovie())

	// Run
	if err := sv.Run(":8080"); err != nil {
		panic(err)
	}
}

// package handlers
func CreateMovie() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title" binding:"required"`
		Rating float64 `json:"rating" binding:"required"`
	}
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"message": "invalid request", "data": nil})
			log.Println("log err: ", err)
			return
		}

		// process
		mv := &Movie{
			ID:     lastID + 1,
			Title:  req.Title,
			Rating: req.Rating,
		}
		movies = append(movies, mv)

		lastID++

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"data":    mv,
		})

	}
}

// package service
type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}

var movies = []*Movie{}
var lastID = 0
