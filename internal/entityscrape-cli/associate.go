package cli

import (
	"github.com/ndabAP/assocentity/v2"
	prose "gopkg.in/jdkato/prose.v2"
)

const (
	unicodeSmallA   = 97
	unicodeSmallZ   = 122
	unicodeCapitalA = 65
	unicodeCapitalZ = 90
)

func weighting(text string, entity []string) (weighting map[string]float64, err error) {
	weighting, err = assocentity.Make(text, entity, func(text string) ([]string, error) {
		tokenizedText, err := tokenizer(text)
		if err != nil {
			return nil, err
		}

		return tokenizedText, nil
	})

	if err != nil {
		return nil, err
	}

	return weighting, nil
}

func tokenizer(text string) ([]string, error) {
	document, err := prose.NewDocument(string(text))
	if err != nil {
		return nil, err
	}

	var tokenizedText []string
	for _, token := range document.Tokens() {
		ok := true
		for _, r := range token.Text {
			// Only allow latin alphabet
			switch {
			case r < unicodeCapitalA:
				ok = false
				break
			case r > unicodeSmallZ:
				ok = false
				break
			case r > unicodeCapitalZ && r < unicodeSmallA:
				ok = false
			}
		}

		if ok {
			tokenizedText = append(tokenizedText, token.Text)
		}
	}

	return tokenizedText, nil
}

func keepAdjectives(weighting map[string]float64) map[string]float64 {
	for word := range weighting {
		if !isInSlice(word, adjectives) {
			delete(weighting, word)
		}
	}

	return weighting
}

func isInSlice(el string, slice []string) bool {
	for _, e := range slice {
		if e == el {
			return true
		}
	}

	return false
}
