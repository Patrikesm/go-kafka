package entity

import "github.com/google/uuid"

type ProductsRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
	FindId(user *User, id string) (*User, error)
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

//construção de entidade
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
