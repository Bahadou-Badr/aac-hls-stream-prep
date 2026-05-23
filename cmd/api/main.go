package main

import (
	"log"
	"net/http"

	"aac-hls-stream-prep/internal/api"
	"aac-hls-stream-prep/internal/hls"
	"aac-hls-stream-prep/internal/storage"
	"aac-hls-stream-prep/internal/track"
	"aac-hls-stream-prep/internal/transcoder"
	"aac-hls-stream-prep/internal/worker"
)

func main() {
	// Repository
	repo := track.NewInMemoryRepository()

	storage := storage.NewLocalStorage("./storage")

	// Create transcoder
	transcoder := transcoder.NewFFmpegTranscoder()

	// Create packager
	packager := hls.NewPackager()

	// Service
	service := track.NewService(repo, storage, transcoder, packager)

	// Worker Pool
	pool := worker.NewPool(100)

	for i := 1; i <= 3; i++ {
		w := worker.NewWorker(
			i,
			service,
			pool.Jobs,
		)

		w.Start()
	}
	// Handler
	handler := api.NewHandler(service, pool)

	// Router
	router := api.NewRouter(handler)

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
