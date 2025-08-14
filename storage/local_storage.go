package storage

import (
	"os"
	"path/filepath"
)

type LocalStorage struct{ base string }

func NewLocalStorage(base string) *LocalStorage {
	_ = os.MkdirAll(base, 0755)
	return &LocalStorage{base: base}
}

func (s *LocalStorage) Save(filename string, data []byte) (string, error) {
	full := filepath.Join(s.base, filename)
	if err := os.WriteFile(full, data, 0644); err != nil {
		return "", err
	}
	// public path served by Gin Static: /uploads/<filename>
	return "/uploads/" + filename, nil
}
