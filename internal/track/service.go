package track

import (
	"aac-hls-stream-prep/internal/hls"
	"aac-hls-stream-prep/internal/storage"
	"aac-hls-stream-prep/internal/transcoder"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

type Service struct {
	repo       Repository
	storage    *storage.LocalStorage
	transcoder *transcoder.FFmpegTranscoder
	packager   *hls.Packager
}

var bitrateLadder = []string{
	"96k",
	"160k",
}

func NewService(repo Repository, storage *storage.LocalStorage, transcoder *transcoder.FFmpegTranscoder, packager *hls.Packager) *Service {
	return &Service{
		repo:       repo,
		storage:    storage,
		transcoder: transcoder,
		packager:   packager,
	}
}

func (s *Service) CreateTrack(id string) (*Track, error) {
	track := &Track{
		ID:     id,
		Status: StatusUploaded,
	}

	err := s.repo.Save(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (s *Service) GetTrack(id string) (*Track, error) {
	return s.repo.GetByID(id)
}

func (s *Service) PrepareTrackUpload(id string, fileReader io.Reader, filename string) (*Track, error) {

	originalPath, err := s.storage.Save(id, fileReader, filename)
	if err != nil {
		return nil, err
	}

	track := &Track{
		ID:       id,
		Status:   StatusUploaded,
		FilePath: originalPath,
		Variants: []AudioVariant{},
	}

	err = s.repo.Save(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (s *Service) ProcessTrack(id string, originalPath string) error {

	track, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	track.Status = StatusProcessing
	if err := s.repo.Save(track); err != nil {
		return err
	}

	var variants []AudioVariant
	trackDir := filepath.Dir(originalPath)

	for _, bitrate := range bitrateLadder {

		outputPath := filepath.Join(trackDir, fmt.Sprintf("track_%s.m4a", bitrate))

		err := s.transcoder.TranscodeToAAC(originalPath, outputPath, bitrate)
		if err != nil {
			return s.failTrack(track, err)
		}

		hlsDir := filepath.Join(trackDir, "hls", bitrate)

		err = s.packager.PackageToHLS(outputPath, hlsDir)
		if err != nil {
			return s.failTrack(track, err)
		}

		variants = append(variants, AudioVariant{
			Bitrate: bitrate,
			Path:    outputPath,
			HLSPath: filepath.Join(hlsDir, "playlist.m3u8"),
			PlaylistURL: fmt.Sprintf(
				"/streams/tracks/%s/hls/%s/playlist.m3u8",
				id,
				bitrate,
			),
		})
	}

	track.Status = StatusReady
	track.Variants = variants

	return s.repo.Save(track)
}

func (s *Service) failTrack(track *Track, cause error) error {
	track.Status = StatusFailed
	if err := s.repo.Save(track); err != nil {
		return errors.Join(cause, err)
	}

	return cause
}
