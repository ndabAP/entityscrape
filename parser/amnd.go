package parser

import (
	"bytes"
	"encoding/json"
	"io"
)

// AMND parses "Adverse Media News Dataset".
func AMND(r io.Reader, texts chan []byte) chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)

		decoder := json.NewDecoder(r)
		if _, err := decoder.Token(); err != nil {
			errs <- err
			return
		}

		var text, lang json.RawMessage
		for decoder.More() {
			token, err := decoder.Token()
			if err != nil {
				errs <- err
				return
			}

			// Key
			k, ok := token.(string)
			if !ok {
				continue
			}

			// Value
			var v json.RawMessage
			if err = decoder.Decode(&v); err != nil {
				return
			}

			switch k {
			case "language":
				lang = v
			case "text":
				text = v
			}
		}

		// Validate
		if !bytes.Equal(lang, []byte("english")) {
			errs <- ErrUnsupportedLang
			return
		}
		if len(text) < 15 {
			errs <- ErrTextTooShort
			return
		}

		// Normalize
		text = bytes.ReplaceAll(text, []byte("\\n"), []byte(" "))
		text = bytes.ReplaceAll(text, []byte("\n"), []byte(" "))
		text = bytes.ReplaceAll(text, []byte("\t"), []byte(" "))
		text = bytes.ReplaceAll(text, []byte("\\u"), []byte(" "))

		texts <- text
	}()

	return errs
}
