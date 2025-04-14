package routes

import (
	"io"
	"net/http"

	"gateway/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Auth)

	// Routes
	r.Get("/orders", handleOrders)
	r.Get("/inventory", handleInventory)

	return r
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://order-service:8081/orders")
	if err != nil {
		http.Error(w, "Failed to reach Order Service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func handleInventory(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://inventory-service:8082/inventory")
	if err != nil {
		http.Error(w, "Failed to reach Inventory Service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
