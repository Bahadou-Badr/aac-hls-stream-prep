package track

import (
	"aac-hls-stream-prep/internal/storage"
	"aac-hls-stream-prep/internal/transcoder"
	"fmt"
	"io"
)

type Service struct {
	repo       Repository
	storage    *storage.LocalStorage
	transcoder *transcoder.FFmpegTranscoder
}

var bitrateLadder = []string{
	"96k",
	"160k",
}

func NewService(repo Repository, storage *storage.LocalStorage, transcoder *transcoder.FFmpegTranscoder) *Service {
	return &Service{
		repo:       repo,
		storage:    storage,
		transcoder: transcoder,
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

func (s *Service) UploadTrack(id string, fileReader io.Reader, filename string) (*Track, error) {
	// Save original upload
	originalPath, err := s.storage.Save(id, fileReader, filename)
	if err != nil {
		return nil, err
	}

	var variants []AudioVariant

	// Generate bitrate ladder
	for _, bitrate := range bitrateLadder {
		outputPath := fmt.Sprintf(
			"./storage/tracks/%s/track_%s.m4a",
			id,
			bitrate,
		)

		err := s.transcoder.TranscodeToAAC(
			originalPath,
			outputPath,
			bitrate,
		)
		if err != nil {
			return nil, err
		}

		variants = append(variants, AudioVariant{
			Bitrate: bitrate,
			Path:    outputPath,
		})
	}

	track := &Track{
		ID:       id,
		Status:   StatusReady,
		FilePath: originalPath,
		Variants: variants,
	}

	err = s.repo.Save(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}
