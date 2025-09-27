package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage[T any] struct {
	FileName string
}

func New[T any](fileName string) *Storage[T] {
	// Ensure the directory exists
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		// If we can't create the directory, we'll let the save operation fail later
	}

	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
