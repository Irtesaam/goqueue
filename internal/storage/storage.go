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

	storage := &Storage[T]{FileName: fileName}

	// Check if this is a fresh installation (directory was just created)
	// If todos.json doesn't exist, we'll initialize it as empty on first load
	return storage
}

// LoadOrInitialize loads data or initializes with empty if file doesn't exist
func (s *Storage[T]) LoadOrInitialize(data *T) error {
	if _, err := os.Stat(s.FileName); os.IsNotExist(err) {
		// File doesn't exist, initialize with empty and save
		// For slices, we need to create an empty slice, not nil
		return os.WriteFile(s.FileName, []byte("[]"), 0644)
	}

	// File exists, load it normally
	return s.Load(data)
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
		// If file doesn't exist, initialize with default value
		if os.IsNotExist(err) {
			return nil // Keep the zero value of T
		}
		return err
	}

	return json.Unmarshal(fileData, data)
}

// InitializeEmpty removes existing file and creates a new empty one with empty array
func (s *Storage[T]) InitializeEmpty() error {
	// Remove the file if it exists
	if _, err := os.Stat(s.FileName); err == nil {
		if err := os.Remove(s.FileName); err != nil {
			return err
		}
	}

	// Create empty array and save it
	var emptyData T
	return s.Save(emptyData)
}
