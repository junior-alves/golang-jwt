package repository

import "github.com/junior-alves/go-test/internal/domain/entity"

type MemoryProductRepository struct {
	products map[string]*entity.Product
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[string]*entity.Product),
	}
}

func (m *MemoryProductRepository) Create(product *entity.Product) bool {
	m.products[product.Id] = product
	return true
}

func (m *MemoryProductRepository) GetAll() ([]*entity.Product, error) {
	products := make([]*entity.Product, len(m.products))

	for _, p := range m.products {
		products = append(products, p)
	}

	return products, nil
}

func (m *MemoryProductRepository) GetProduct(id string) *entity.Product {
	return m.products[id]
}
