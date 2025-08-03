package parser

import (
	"encoding/csv"
	"io"
)

func DS(r io.Reader) (texts []string, err error) {
	c := csv.NewReader(r)
	c.TrimLeadingSpace = true

	c.Read()

	d, err := c.ReadAll()
	if err != nil {
		return texts, err
	}
	for _, r := range d[1:] {
		texts = append(texts, r[9])
	}
	return texts, nil
}
