package services

import (
	"inventory/internal/models"
	"inventory/internal/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}
