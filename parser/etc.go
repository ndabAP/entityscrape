package parser

import (
	"bufio"
	"io"
	"strings"
)

func Etc(r io.Reader) (texts []string, err error) {
	scanner := bufio.NewScanner(r)
	var txt strings.Builder
	for scanner.Scan() {
		txt.WriteString(scanner.Text())
		txt.WriteString(".")
	}
	texts = []string{txt.String()}

	err = scanner.Err()
	return texts, err
}
