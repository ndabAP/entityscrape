package parser

import (
	"encoding/csv"
	"io"
)

func DS(r io.Reader, c chan []byte) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)

		reader := csv.NewReader(r)
		reader.TrimLeadingSpace = true

		records, err := reader.ReadAll()
		if err != nil {
			errs <- err
			return
		}

		for _, r := range records[1:] {
			c <- []byte(r[9])
		}
	}()

	return errs
}
