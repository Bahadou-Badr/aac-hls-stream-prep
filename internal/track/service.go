package track

import (
	"aac-hls-stream-prep/internal/storage"
	"io"
)

type Service struct {
	repo    Repository
	storage *storage.LocalStorage
}

func NewService(repo Repository, storage *storage.LocalStorage) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
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
	path, err := s.storage.Save(id, fileReader, filename)
	if err != nil {
		return nil, err
	}

	track := &Track{
		ID:       id,
		Status:   StatusUploaded,
		FilePath: path,
	}

	err = s.repo.Save(track)
	if err != nil {
		return nil, err
	}

	return track, nil
}
