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

		record, err := reader.Read()
		if err != nil {
			errs <- err
			return
		}

		c <- []byte(record[9])
	}()

	return errs
}
