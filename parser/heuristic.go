package parser

import (
	"bufio"
	"bytes"
)

type ApplyHeuristic func(text, entity []byte) ([]byte, error)

var DotHeuristic ApplyHeuristic = func(text, entity []byte) ([]byte, error) {
	scanner := bufio.NewScanner(bytes.NewReader(text))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if i := bytes.IndexByte(data, '.'); i >= 0 {
			return i + 1, data[:i], nil
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	t := make([]byte, 0)
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), entity) {
			t = append(t, scanner.Bytes()...)
		}
	}
	return t, scanner.Err()
}
