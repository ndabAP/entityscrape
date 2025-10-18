package parser

import (
	"bufio"
	"bytes"
	"io"
)

func CRFTC(r io.Reader, c chan []byte) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)

		scanner := bufio.NewScanner(r)
		scanner.Scan() // Skip header
		for scanner.Scan() {
			spl := bytes.Split(scanner.Bytes(), []byte("|"))
			if len(spl) < 2 {
				continue
			}

			text := spl[1]

			// Normalize
			text = bytes.ReplaceAll(text, []byte(" e "), []byte(" "))

			c <- text
		}
		if err := scanner.Err(); err != nil {
			errs <- err
		}
	}()

	return errs
}
