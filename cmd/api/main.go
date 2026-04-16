package main

import (
	"log"
	"net/http"

	"aac-hls-stream-prep/internal/api"
	"aac-hls-stream-prep/internal/storage"
	"aac-hls-stream-prep/internal/track"
)

func main() {
	// Repository
	repo := track.NewInMemoryRepository()

	storage := storage.NewLocalStorage("./storage")

	// Service
	service := track.NewService(repo, storage)

	// Handler
	handler := api.NewHandler(service)

	// Router
	router := api.NewRouter(handler)

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
