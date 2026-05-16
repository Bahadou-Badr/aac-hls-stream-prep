package track

type Status string

type AudioVariant struct {
	Bitrate string `json:"bitrate"`
	Path    string `json:"path"`
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
