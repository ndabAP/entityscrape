package sbd

import "golang.org/x/text/language"

// Tokenizer writes tokenized sentences to channel c using sentence boundary
// disambiguation.
func Tokenize(lang language.Tag, text string, c chan string) {
	switch lang {
	case language.German:
		for _, s := range tokenizer.de.Tokenize(text) {
			c <- s.Text
		}

	case language.English:
		for _, s := range tokenizer.en.Tokenize(text) {
			c <- s.Text
		}

	case language.Russian:
		for _, s := range tokenizer.ru.Segment(text) {
			c <- s
		}

	default:
		panic("unsupported language")
	}
}
