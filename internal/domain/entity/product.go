package entity

import "github.com/google/uuid"

type Product struct {
	Id    string
	Name  string
	Price int
}

func NewProduct(name string, price int) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
