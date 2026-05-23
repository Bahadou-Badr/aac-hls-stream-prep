package hls

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Packager struct{}

func NewPackager() *Packager {
	return &Packager{}
}

func (p *Packager) PackageToHLS(inputPath string, outputDir string) error {

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	playlistPath := filepath.Join(outputDir, "playlist.m3u8")

	cmd := exec.Command(
		"ffmpeg",
		"-i", inputPath,

		// No re-encoding
		"-c:a", "copy",

		// HLS muxer
		"-f", "hls",

		// CMAF fragmented MP4
		"-hls_segment_type", "fmp4",

		// VOD
		"-hls_playlist_type", "vod",

		// Segment duration
		"-hls_time", "6",

		// Segment naming
		"-hls_segment_filename",
		filepath.Join(outputDir, "segment_%03d.m4s"),

		// Explicit init segment path
		"-hls_fmp4_init_filename",
		filepath.Join(outputDir, "init.mp4"),

		playlistPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"hls packaging error: %s",
			string(output),
		)
	}

	return nil
}
