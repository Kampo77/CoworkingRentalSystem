package controllers

import (
	"net/http"
	"order/internal/models"
	"order/internal/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.service.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order, err := oc.service.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")

	var orderUpdate struct {
		Status string `json:"status"`
	}

	if err := c.BindJSON(&orderUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.service.UpdateOrderStatus(id, orderUpdate.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}

func (oc *OrderController) GetOrdersByUserID(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	orders, err := oc.service.GetOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
