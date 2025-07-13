package parser

import (
	"encoding/json"
	"io"
	"strings"
)

// AMND parses "Adverse Media News Dataset".
func AMND(r io.Reader) (text []string, err error) {
	type data struct {
		Text string `json:"text"`
	}

	var d data
	if err = json.NewDecoder(r).Decode(&d); err != nil {
		return
	}
	text = []string{d.Text}

	// Validate
	if len(text) < 15 {
		return []string{}, ErrTextTooShort
	}

	// Replace
	text[0] = strings.ReplaceAll(text[0], "\n", " ")
	text[0] = strings.ReplaceAll(text[0], "\t", " ")

	return
}
