package track

import (
	"aac-hls-stream-prep/internal/hls"
	"aac-hls-stream-prep/internal/storage"
	"aac-hls-stream-prep/internal/transcoder"
	"fmt"
	"io"
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

	var variants []AudioVariant

	for _, bitrate := range bitrateLadder {

		outputPath := fmt.Sprintf(
			"./storage/tracks/%s/track_%s.m4a",
			id,
			bitrate,
		)

		err := s.transcoder.TranscodeToAAC(originalPath, outputPath, bitrate)
		if err != nil {
			track.Status = StatusFailed
			return err
		}

		hlsDir := fmt.Sprintf("./storage/tracks/%s/hls/%s", id, bitrate)

		err = s.packager.PackageToHLS(outputPath, hlsDir)
		if err != nil {
			track.Status = StatusFailed
			return err
		}

		variants = append(variants, AudioVariant{
			Bitrate: bitrate,
			Path:    outputPath,
			HLSPath: fmt.Sprintf(
				"%s/playlist.m3u8",
				hlsDir,
			),
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
