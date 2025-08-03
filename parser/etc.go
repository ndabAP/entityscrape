package parser

import (
	"bufio"
	"bytes"
	"io"
)

func Etc(r io.Reader, texts chan []byte) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)

		scanner := bufio.NewScanner(r)
		var buf bytes.Buffer
		for scanner.Scan() {
			buf.Write(scanner.Bytes())
			buf.WriteRune('.')
		}

		texts <- buf.Bytes()
		errs <- scanner.Err()
	}()

	return errs
}
