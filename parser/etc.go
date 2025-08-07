package parser

import (
	"bufio"
	"bytes"
	"io"
)

func Etc(r io.Reader, c chan []byte) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)

		scanner := bufio.NewScanner(r)
		var buf bytes.Buffer
		for scanner.Scan() {
			buf.Write(scanner.Bytes())
			buf.WriteRune('.')
		}

		c <- buf.Bytes()
		errs <- scanner.Err()
	}()

	return errs
}
