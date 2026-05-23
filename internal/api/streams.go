package api

type StreamVariant struct {
	Bitrate string `json:"bitrate"`
	URL     string `json:"url"`
}

type StreamsResponse struct {
	TrackID string          `json:"track_id"`
	Status  string          `json:"status"`
	Streams []StreamVariant `json:"streams"`
}
