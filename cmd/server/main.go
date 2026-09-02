package main

import (
	"log"
	"net/http"

	"github.com/Nesdann/webhook-dispatcher/internal/api"
)

func main() {
	http.HandleFunc("/api/v1/event", api.Handler)
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}