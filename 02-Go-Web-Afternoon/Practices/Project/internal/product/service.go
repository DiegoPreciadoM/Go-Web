package product

import (
	"errors"

	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchByPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	l := s.r.GetAll()
	return l, nil
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) SearchByPriceGt(price float64) ([]domain.Product, error) {
	l := s.r.SearchByPriceGt(price)
	if len(l) == 0 {
		return []domain.Product{}, errors.New("No products found")
	}
	return l, nil
}

func (s *service) Create(p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(id int, u domain.Product) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.Code_value != "" {
		p.Code_value = u.Code_value
	}
	if u.Expiration != "" {
		p.Expiration = u.Expiration
	}
	if u.Quantity > 0 {
		p.Quantity = u.Quantity
	}
	if u.Price > 0 {
		p.Price = u.Price
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil

}
