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
func GPSC(r io.Reader) ([]string, error) {
	texts := make([]string, 0)

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
		return texts, err
	}
	for _, t := range coll.Texts {
		text := fmt.Sprintf("%s %s", t.Anrede, t.Rohtext)

		if len(text) < 3 {
			slog.Debug("skipping short text")
			continue
		}
		text = html.UnescapeString(text)
		text = spaceregex.ReplaceAllString(text, " ")

		texts = append(texts, text)
	}

	return texts, nil
}
