package parser

import (
	"encoding/json"
	"io"
	"strings"
)

// AMND parses "Adverse Media News Dataset".
func AMND(f io.Reader) (text string, err error) {
	type data struct {
		Text string `json:"text"`
	}

	var d data
	if err = json.NewDecoder(f).Decode(&d); err != nil {
		return
	}
	text = d.Text

	// Validate
	if len(text) < 15 {
		return "", ErrTextTooShort
	}

	// Replace
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	return
}
