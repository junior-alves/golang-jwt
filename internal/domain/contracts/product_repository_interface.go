package contracts

import "github.com/junior-alves/go-test/internal/domain/entity"

type ProductRepository interface {
	Create(product *entity.Product) bool
	GetAll() ([]*entity.Product, error)
	GetProduct(id string) *entity.Product
}
