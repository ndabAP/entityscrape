package parser

import (
	"encoding/csv"
	"io"
)

func DS(r io.Reader) (text []string, err error) {
	c := csv.NewReader(r)
	c.TrimLeadingSpace = true

	d, err := c.ReadAll()
	if err != nil {
		return text, err
	}
	for _, r := range d[1:] {
		text = append(text, r[9])
	}
	return text, nil
}
