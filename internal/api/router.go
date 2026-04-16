package api

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTrack(w, r)
		case http.MethodGet:
			handler.GetTrack(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tracks/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.UploadTrack(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	return mux
}
