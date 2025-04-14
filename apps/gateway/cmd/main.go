package main

import (
	"gateway/internal/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRouter()

	log.Println("🚀 Gateway started on port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
