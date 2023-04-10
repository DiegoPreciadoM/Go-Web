package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func main() {

	router := gin.Default()

	router.POST("/saludo", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}

		myReader := strings.NewReader(string(body))
		myDecode := json.NewDecoder(myReader)

		var info Info

		if err := myDecode.Decode(&info); err != nil {
			log.Fatal(err)
		}

		response := "Hola " + info.Name + " " + info.Lastname

		c.Data(http.StatusAccepted, "application/json; charset=utf-8", []byte(response))
	})

	router.Run(":8080")

}
