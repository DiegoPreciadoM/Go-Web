package product

import (
	"errors"

	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/internal/domain"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchByPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	list []domain.Product
}

func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

func (r *repository) GetAll() []domain.Product {
	return r.list
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("Product not found")
}

func (r *repository) SearchByPriceGt(price float64) []domain.Product {
	var products []domain.Product

	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}

	return products
}

func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.Code_value) {
		return domain.Product{}, errors.New("Code value already exist")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

func (r *repository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.Code_value == codeValue {
			return false
		}
	}
	return true
}

func (r *repository) Update(id int, p domain.Product) (domain.Product, error) {
	for i, product := range r.list {
		if product.Id == id {
			if !r.validateCodeValue(product.Code_value) && product.Code_value != p.Code_value {
				return domain.Product{}, errors.New("Code value already exist")
			}
			r.list[i] = p
			return p, nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}

func (r *repository) Delete(id int) error {
	for i, product := range r.list {
		if product.Id == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			return nil
		}
	}
	return errors.New("Product not found")
}
