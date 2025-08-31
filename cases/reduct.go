package cases

import (
	"bufio"
	"bytes"
	"errors"
	"unicode"
	"unicode/utf8"
)

var errEntityNotFound = errors.New("entity not found")

// reduct perfoms a fuzzy search for the provided entity and removes sentences
// defined through [unicode.Sentence_Terminal] which don't contain the entity.
func (study study[samples, aggregated]) reduct(text []byte, entity []string) ([]byte, error) {
	var (
		// Cached delimiter
		delim rune
		size  int
	)

	scanner := bufio.NewScanner(bytes.NewReader(text))
	scanner.Split(func(txt []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(txt) == 0 {
			return 0, nil, nil
		}

		// Find the index of the first byte of a rune that is a sentence
		// terminal.
		i := bytes.IndexFunc(txt, func(r rune) bool {
			return unicode.Is(unicode.Sentence_Terminal, r)
		})
		switch i {
		// Not found.
		case -1:
			if atEOF && len(txt) > 0 {
				return len(txt), txt, nil
			}
			return 0, nil, nil

		// We found a sentence terminal. We need to know its size to advance
		// past it.
		default:
			delim, size = utf8.DecodeRune(txt[i:])
			return i + size, txt[:i], nil
		}
	})

	var (
		buf bytes.Buffer

		// Filters
		prot = []byte("http")
	)
	for scanner.Scan() {
		txt := scanner.Bytes()

		// Entity
		var n int
		for _, e := range entity {
			if r, _ := utf8.DecodeRuneInString(e); r == delim {
				n++
			}
			if bytes.Contains(txt, []byte(e)) {
				n++
			}
		}
		if n == 0 {
			continue
		}
		// Filters
		if bytes.Contains(txt, prot) {
			continue
		}
		// Size
		if len(txt) < 3 {
			continue
		}
		n = 0
		for _, r := range scanner.Text() {
			if unicode.IsSpace(r) {
				n++
			}
		}
		if n == 0 {
			continue
		}

		buf.Write(txt)
		// Re-add any terminal.
		buf.WriteRune(delim)
	}

	b := buf.Bytes()
	if len(b) == 0 {
		return []byte{}, errEntityNotFound
	}
	return bytes.TrimSpace(b), nil // Split never returns an error.
}
