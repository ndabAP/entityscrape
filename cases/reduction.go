package cases

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

func reduce(text, entity string) (string, error) {
	var (
		delim rune
		size  int
	)
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the first byte of a rune that is a sentence
		// terminal.
		i := bytes.IndexFunc(data, func(r rune) bool {
			return unicode.Is(unicode.Sentence_Terminal, r)
		})
		switch i {
		// Not found.
		case -1:
			if atEOF && len(data) > 0 {
				return len(data), data, nil
			}
			return 0, nil, nil

		// We found a sentence terminal. We need to know its size to advance
		// past it.
		default:
			delim, size = utf8.DecodeRune(data[i:])
			return i + size, data[:i], nil
		}
	})

	var t strings.Builder
	for scanner.Scan() {
		if !bytes.Contains(scanner.Bytes(), []byte(entity)) {
			continue
		}

		if _, err := t.Write(scanner.Bytes()); err != nil {
			return "", err
		}

		// Re-add any terminal.
		t.WriteRune(delim)
	}

	return t.String(), scanner.Err()
}
