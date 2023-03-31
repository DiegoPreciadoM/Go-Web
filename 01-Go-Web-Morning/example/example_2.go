package main

import (
	"encoding/json"
	"os"
)

func main() {
	var err error

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

	standardEncoder := json.NewEncoder(os.Stdout)

	if err = standardEncoder.Encode(data); err != nil {
		panic(err)
	}
}
