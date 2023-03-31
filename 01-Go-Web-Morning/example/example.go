package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	IsMarried   bool     `json:"isMarried"`
	ShopingList []string `json:"-"`
}

func main() {
	data := map[string]any{
		"name":      "John",
		"age":       30,
		"isMarried": true,
		"shopList": []string{
			"milk",
			"apple",
			"coffee",
		},
	}

	dataAsJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data as map: %+v", data)
	fmt.Printf("Data as Json %s", dataAsJson)

	user := User{
		Name:      "John",
		Age:       30,
		IsMarried: true,
		ShopingList: []string{
			"milk",
			"apple",
			"coffee",
		},
	}

	structAsJson, err_s := json.Marshal(user)
	if err_s != nil {
		panic(err_s)
	}

	fmt.Printf("User as struct: %+v", user)
	fmt.Printf("User as Json %s", structAsJson)

	var userMapFromJson map[string]any

	if err = json.Unmarshal(structAsJson, &userMapFromJson); err != nil {
		panic(err)
	}

	fmt.Println("User as Map Unmarshal ", userMapFromJson)
}
