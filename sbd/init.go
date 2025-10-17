package sbd

import (
	"os"

	"github.com/neurosnap/sentences"
	"github.com/neurosnap/sentences/english"
	"github.com/sentencizer/sentencizer"
)

var tokenizer struct {
	de, en sentences.SentenceTokenizer
	ru     sentencizer.Segmenter
}

func init() {
	// German
	{
		b, err := os.ReadFile("./third_party/neurosnap/sentences/data/german.json")
		if err != nil {
			panic(err.Error())
		}
		traning, err := sentences.LoadTraining(b)
		if err != nil {
			panic(err.Error())
		}
		tokenizer.de = sentences.NewSentenceTokenizer(traning)
	}
	// English
	{
		t, err := english.NewSentenceTokenizer(nil)
		if err != nil {
			panic(err)
		}
		tokenizer.en = t
	}
	// Russian
	{
		tokenizer.ru = sentencizer.NewSegmenter("ru")
	}
}
