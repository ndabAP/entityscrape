package cases

import (
	"io/fs"
	"log/slog"
	"path/filepath"
)

func WalkCorpus(corpus string, fn func(filename string) error) error {
	root := filepath.Join("corpus", corpus)
	err := fs.WalkDir(Corpus, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".gitignore" {
			return nil
		}

		slog.Debug("processing file", "filename", path)
		return fn(path)
	})
	return err
}
