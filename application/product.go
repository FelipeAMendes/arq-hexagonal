package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	Id     string
	Name   string
	Price  float64
	Status string
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == DISABLED {
		return false, nil
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("The price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	p.Status = DISABLED
	return nil
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		Id:     id,
		Name:   name,
		Price:  price,
		Status: DISABLED,
	}
}
