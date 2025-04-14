package services

import (
	"order/internal/models"
	"order/internal/repositories"
)

type OrderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	order.Status = "pending"
	return s.repo.Create(order)
}

func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *OrderService) UpdateOrderStatus(id string, status string) error {
	// Можно добавить валидацию статуса здесь
	// например, проверить, что status принадлежит списку допустимых значений
	return s.repo.UpdateStatus(id, status)
}

func (s *OrderService) GetOrdersByUserID(userID string) ([]models.Order, error) {
	return s.repo.GetByUserID(userID)
}
