package track

type Status string

const (
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusReady      Status = "ready"
	StatusFailed     Status = "failed"
)

type Track struct {
	ID       string `json:"id"`
	Status   Status `json:"status"`
	FilePath string `json:"file_path"`
	AACPath  string `json:"aac_path"`
}
