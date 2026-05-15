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

	// AAC output path
	aacOutputPath := fmt.Sprintf(
		"./storage/tracks/%s/track_160k.m4a",
		id,
	)

	// Transcode to AAC
	err = s.transcoder.TranscodeToAAC(
		originalPath,
		aacOutputPath,
		"160k",
	)
	if err != nil {
		return nil, err
	}

	track := &Track{
		ID:       id,
		Status:   StatusReady,
		FilePath: originalPath,
		AACPath:  aacOutputPath,
	}

	err = s.repo.Save(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}
