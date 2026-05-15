package transcoder

import (
	"fmt"
	"os/exec"
)

type FFmpegTranscoder struct{}

func NewFFmpegTranscoder() *FFmpegTranscoder {
	return &FFmpegTranscoder{}
}

func (t *FFmpegTranscoder) TranscodeToAAC(inputPath, outputPath string, bitrate string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputPath,
		"-c:a", "aac",
		"-b:a", bitrate,
		outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ffmpeg error: %s", string(output))
	}

	return nil
}
