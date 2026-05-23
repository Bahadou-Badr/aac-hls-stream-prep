package api

import (
	"encoding/json"
	"net/http"

	"aac-hls-stream-prep/internal/track"
	"aac-hls-stream-prep/internal/worker"
	"aac-hls-stream-prep/pkg/id"
)

type Handler struct {
	service *track.Service
	pool    *worker.Pool
}

func NewHandler(service *track.Service, pool *worker.Pool) *Handler {
	return &Handler{service: service, pool: pool}
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

	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	trackID := id.New()

	t, err := h.service.PrepareTrackUpload(
		trackID,
		file,
		header.Filename,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Queue background job
	h.pool.Jobs <- worker.Job{
		TrackID:      t.ID,
		OriginalPath: t.FilePath,
	}

	writeJSON(w, http.StatusAccepted, t)
}

func (h *Handler) GetStreams(w http.ResponseWriter, r *http.Request) {

	trackID := r.URL.Query().Get("id")

	if trackID == "" {
		http.Error(w, "track id required", http.StatusBadRequest)
		return
	}

	track, err := h.service.GetTrack(trackID)
	if err != nil {
		http.Error(w, "track not found", http.StatusNotFound)
		return
	}

	var streams []StreamVariant

	for _, variant := range track.Variants {

		streams = append(streams, StreamVariant{
			Bitrate: variant.Bitrate,
			URL:     variant.PlaylistURL,
		})
	}

	response := StreamsResponse{
		TrackID: track.ID,
		Status:  string(track.Status),
		Streams: streams,
	}

	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
