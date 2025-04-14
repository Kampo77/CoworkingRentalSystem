package main

import (
	"log"
	"order/internal/controllers"
	"order/internal/repositories"
	"order/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var orderRepo repositories.OrderRepository
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders/:id", orderController.GetOrderByID)
	router.PATCH("/orders/:id", orderController.UpdateOrderStatus)
	router.GET("/orders", orderController.GetOrdersByUserID)

	log.Fatal(router.Run(":8082"))
}
