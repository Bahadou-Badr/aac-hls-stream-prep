package api

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	// Tracks
	mux.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTrack(w, r)

		case http.MethodGet:
			handler.GetTrack(w, r)

		default:
			http.Error(
				w,
				"method not allowed",
				http.StatusMethodNotAllowed,
			)
		}
	})

	// Upload
	mux.HandleFunc("/tracks/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.UploadTrack(w, r)
			return
		}

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	})

	// Streams API
	mux.HandleFunc("/tracks/streams", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetStreams(w, r)
			return
		}

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	})

	// Static HLS/CMAF delivery
	fs := http.FileServer(http.Dir("./storage"))

	mux.Handle(
		"/streams/",
		http.StripPrefix("/streams/", fs),
	)

	return mux
}
