package cases

import (
	"io/fs"
	"path/filepath"
)

// corpusDir is the absolute directory to the corpus.
var corpusDir string

func SetCorpusRootDir(base string) {
	corpusDir = filepath.Join(base, "corpus")
}

func GetCorpusRootDir() string {
	return corpusDir
}

func WalkCorpus(corpus string, fn func(filename string) error) error {
	root := filepath.Join(corpusDir, corpus)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(path) {
		case ".gitignore", ".gitkeep":
			return nil
		}
		return fn(path)
	})
	return err
}
