package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EXMAPLE 1 - JSON.MARSHAL(v interface{}) []byte, err

type Product struct {
	Name      string `json:"Name"`
	Price     int    `json:"Price"`
	Published bool   `json:"Published"`
}

// func main() {

// 	p := Product{
// 		Name:      "MackBook Pro",
// 		Price:     1500,
// 		Published: true,
// 	}

// 	jsonData, err := json.Marshal(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(jsonData))
// }

// EXAMPLE 2 - JSON.UNMARSHAL(Array bytes([]byte), puntero de cualquier tipo (v interface{})) err

// func main() {
// 	jsonData := `{"Name":"MackBook Pro M1","Price":1700,"Published":true}`

// 	var p Product

// 	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(p.Name)
// 	fmt.Println(p.Price)
// 	fmt.Println(p.Published)
// }

// EXAMPLE 3 - JSON.ENCODE()

// func main() {

// 	myEncoder := json.NewEncoder(os.Stdout)

// 	data := Product{
// 		Name:      "MackBook Pro M2",
// 		Price:     2000,
// 		Published: true,
// 	}

// 	myEncoder.Encode(data)

// }

// EXAMPLE 4 - JSON.DECODER()

// const jsonStream = `
// 	{"Name":"MackBook Pro M1","Price":1700,"Published":true}
// 	{"Name":"MackBook Pro M2","Price":1800,"Published":false}
// 	{"Name":"MackBook Pro M3","Price":2000,"Published":true}`

// func main() {

// 	myStreaming := strings.NewReader(jsonStream)
// 	myDecoder := json.NewDecoder(myStreaming)

// 	var products []Product
// 	index := 0

// 	for {

// 		var product Product

// 		if err := myDecoder.Decode(&product); err == io.EOF {
// 			break
// 		} else {
// 			fmt.Println(index)
// 			log.Fatal(err)
// 		}

// 		products = append(products, product)

// 		fmt.Println("Datos Recibidos:", products[index].Name, products[index].Price, products[index].Published)
// 		index++

// 	}

// 	fmt.Println(products)

// }

// EXAMPLE 5 - GIN ( CREATING A SERVER )

func main() {
	// Creating a server start
	router := gin.Default()

	// Creating a handler
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})
	})

	// Run server
	router.Run(":8080")
}
