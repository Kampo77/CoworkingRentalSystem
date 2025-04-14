package repositories

import (
	"order/internal/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
	UpdateStatus(id string, status string) error
	GetByUserID(userID string) ([]models.Order, error)
}
