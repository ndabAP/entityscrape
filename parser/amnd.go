package parser

import (
	"bytes"
	"encoding/json"
	"io"
)

// AMND parses "Adverse Media News Dataset".
func AMND(r io.Reader, heuristics ...ApplyHeuristic) (text []string, err error) {
	// TODO: Text is empty
	type data struct {
		Language string          `json:"language"`
		Text     json.RawMessage `json:"text"`
	}

	var d data
	if err = json.NewDecoder(r).Decode(&d); err != nil {
		return
	}

	// Validate
	if len(d.Text) < 15 {
		return []string{}, ErrTextTooShort
	}

	switch d.Language {
	case "english":
	default:
		return []string{}, ErrUnsupportedLang
	}

	// Normalize
	d.Text = bytes.ReplaceAll(d.Text, []byte("\\n"), []byte(" "))
	d.Text = bytes.ReplaceAll(d.Text, []byte("\n"), []byte(" "))
	d.Text = bytes.ReplaceAll(d.Text, []byte("\t"), []byte(" "))
	d.Text = bytes.ReplaceAll(d.Text, []byte("\\u"), []byte(" "))

	for _, h := range heuristics {
	}

	text = []string{d.Text}
	return
}
