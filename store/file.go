package store

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type file struct {
	base string
}

var mu sync.Mutex

func NewFile(base string) (*file, error) {
	mu.Lock()
	defer mu.Unlock()

	f := &file{base: base}
	if err := f.RemoveAll(); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(base, 0o755); err != nil {
		return nil, err
	}

	return f, nil
}

func (s *file) RemoveAll() error {
	return os.RemoveAll(s.base)
}

func (s *file) NewWriter(pref, ext string) (io.WriteCloser, error) {
	filename := filepath.Join(s.base, fmt.Sprintf("%s.%s", pref, ext))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
