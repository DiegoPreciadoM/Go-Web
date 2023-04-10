package Products

import (
	"encoding/json"
	"errors"
	"os"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required"`
	Code_value   string  `json:"code_value" binding:"required"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

func LoadInformation(filename string) ([]Product, error) {
	var products []Product

	file, err := os.Open(filename)
	if err != nil {
		return products, errors.New("Error reading text")
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&products); err != nil {
		return products, errors.New("Error decoding file")
	}

	file.Close()
	return products, nil
}

type SaveProduct struct {
	products []Product
}

func (sv *SaveProduct) GetProducts(filename string) error {
	var err error
	sv.products, err = LoadInformation(filename)
	if err != nil {
		return err
	}

	return nil
}

func (sv *SaveProduct) GetInformationProducts() []Product {
	return sv.products
}

func (sv *SaveProduct) SaveNewProduct(new_product Product) {
	sv.products = append(sv.products, new_product)
}
