package cases

import (
	"bufio"
	"bytes"
	"unicode"
	"unicode/utf8"
)

// reduct perfoms a fuzzy search for the provided entity and removes sentences
// defined in [unicode.Sentence_Terminal] which don't contain the entity.
func reduce(text []byte, entity string) ([]byte, error) {
	var (
		// Cached delimiter
		delim rune
		size  int
	)

	scanner := bufio.NewScanner(bytes.NewReader(text))
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

	var (
		buf bytes.Buffer
		e   = []byte(entity)
	)
	for scanner.Scan() {
		if !bytes.Contains(scanner.Bytes(), e) {
			continue
		}

		if _, err := buf.Write(scanner.Bytes()); err != nil {
			return []byte{}, err
		}

		// Re-add any terminal.
		buf.WriteRune(delim)
	}

	return buf.Bytes(), scanner.Err()
}
