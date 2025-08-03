package parser

import (
	"encoding/csv"
	"io"
)

func DS(r io.Reader, texts chan []byte) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)
		c := csv.NewReader(r)
		c.TrimLeadingSpace = true

		record, err := c.Read()
		if err != nil {
			errs <- err
			return
		}

		texts <- []byte(record[9])
	}()

	return errs
}
