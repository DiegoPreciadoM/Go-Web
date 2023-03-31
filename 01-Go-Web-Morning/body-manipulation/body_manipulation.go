package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func main() {
	var err error

	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		myStreaming := strings.NewReader(string(body))
		myDecoder := json.NewDecoder(myStreaming)

		var user User

		if err := myDecoder.Decode(&user); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"Name":      user.Name,
			"Last Name": user.LastName,
		})
	})

	if err = router.Run(":8080"); err != nil {
		panic(err)
	}
}
