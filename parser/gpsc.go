package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"regexp"

	"golang.org/x/net/html"
)

var spaceregex = regexp.MustCompile(`\s+`)

// GPSC parses "German Political Speeches Corpus".
func GPSC(r io.Reader, c chan []byte) chan error {
	errs := make(chan error, 1)

	go func() {
		defer close(errs)

		type (
			text struct {
				Anrede  string `xml:"anrede,attr"`
				Rohtext string `xml:"rohtext"`
			}
			collection struct {
				Texts []text `xml:"text"`
			}
		)

		coll := collection{}
		if err := xml.NewDecoder(r).Decode(&coll); err != nil {
			errs <- err
			return
		}
		for _, t := range coll.Texts {
			text := fmt.Sprintf("%s %s", t.Anrede, t.Rohtext)

			if len(text) < 15 {
				slog.Debug(ErrTextTooShort.Error())
				continue
			}
			text = html.UnescapeString(text)
			text = spaceregex.ReplaceAllString(text, " ")

			c <- []byte(text)
		}
	}()

	return errs
}
