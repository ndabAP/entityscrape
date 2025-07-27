package parser

import (
	"encoding/json"
	"io"
	"strings"
)

// AMND parses "Adverse Media News Dataset".
func AMND(r io.Reader) (text []string, err error) {
	type data struct {
		Language string `json:"language"`
		Text     string `json:"text"`
	}

	var d data
	if err = json.NewDecoder(r).Decode(&d); err != nil {
		return
	}

	// Validate
	if len(text) < 15 {
		return []string{}, ErrTextTooShort
	}

	text = []string{d.Text}

	switch d.Language {
	case "english":
	default:
		return []string{}, ErrUnsupportedLang
	}

	// Normalize
	text[0] = strings.ReplaceAll(text[0], "\n", " ")
	text[0] = strings.ReplaceAll(text[0], "\t", " ")

	return
}
