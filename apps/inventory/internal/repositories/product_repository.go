package repositories

import (
	"inventory/internal/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
	GetAll() ([]models.Product, error)
}
