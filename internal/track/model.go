package track

type Status string

type AudioVariant struct {
	Bitrate     string `json:"bitrate"`
	Path        string `json:"path"`
	HLSPath     string `json:"hls_path"`
	PlaylistURL string `json:"playlist_url"`
}

const (
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusReady      Status = "ready"
	StatusFailed     Status = "failed"
)

type Track struct {
	ID       string         `json:"id"`
	Status   Status         `json:"status"`
	FilePath string         `json:"file_path"`
	Variants []AudioVariant `json:"variants"`
}
