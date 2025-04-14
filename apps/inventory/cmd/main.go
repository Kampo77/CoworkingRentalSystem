package main

import (
	"inventory/internal/controllers"
	"inventory/internal/repositories"
	"inventory/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var productRepo repositories.ProductRepository
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	router.POST("/products", productController.CreateProduct)
	router.GET("/products/:id", productController.GetProductByID)
	router.PATCH("/products/:id", productController.UpdateProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
	router.GET("/products", productController.GetProducts)

	log.Fatal(router.Run(":8081"))
}
