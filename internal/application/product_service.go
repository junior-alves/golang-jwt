package application

import (
	"errors"

	"github.com/junior-alves/go-test/internal/domain/contracts"
	"github.com/junior-alves/go-test/internal/domain/entity"
)

type ProductService struct {
	productRepository contracts.ProductRepository
}

func NewProductService(repository contracts.ProductRepository) *ProductService {
	return &ProductService{productRepository: repository}
}

func (s *ProductService) CreateProduct(name string, price int) *entity.Product {
	product := entity.NewProduct(name, price)
	s.productRepository.Create(product)
	return product
}

func (s *ProductService) ListProducts() []*entity.Product {
	products, _ := s.productRepository.GetAll()

	if products == nil {
		return make([]*entity.Product, 0)
	}

	return products
}

func (s *ProductService) GetProduct(id string) (*entity.Product, error) {
	product := s.productRepository.GetProduct(id)

	if product == nil {
		return nil, errors.New("Not Found")
	}

	return product, nil
}
