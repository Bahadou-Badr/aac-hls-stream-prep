package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{basePath: basePath}
}

func (s *LocalStorage) Save(trackID string, file io.Reader, filename string) (string, error) {
	dir := filepath.Join(s.basePath, "tracks", trackID)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	//normalize the uploaded filename
	extension := filepath.Ext(filename)

	normalizedFilename := "original" + extension

	filePath := filepath.Join(dir, normalizedFilename)

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
