package api

import (
	"encoding/json"
	"net/http"

	"aac-hls-stream-prep/internal/track"
	"aac-hls-stream-prep/pkg/id"
)

type Handler struct {
	service *track.Service
}

func NewHandler(service *track.Service) *Handler {
	return &Handler{service: service}
}

// POST /tracks
func (h *Handler) CreateTrack(w http.ResponseWriter, r *http.Request) {
	trackID := id.New()

	t, err := h.service.CreateTrack(trackID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

// GET /tracks?id=123 (simple version for Phase 1)
func (h *Handler) GetTrack(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	t, err := h.service.GetTrack(id)
	if err != nil {
		if err == track.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) UploadTrack(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	trackID := id.New()

	t, err := h.service.UploadTrack(trackID, file, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
