package cases

import (
	"io"

	"github.com/ndabAP/entitydebs"
	"github.com/ndabAP/entitydebs/tokenize"
	"golang.org/x/text/language"
)

type (
	Parser   func(io.Reader, chan []byte) chan error
	Analyses struct {
		Entity []string

		Language language.Tag

		Feats         tokenize.Features
		Parser        Parser
		FuzzyMatching bool

		Ext       string
		Filenames []string
	}

	Translate func([]string) ([]string, error)

	Collector[samples any]              func(entitydebs.Frames) samples
	Aggregator[samples, aggregated any] func(samples) aggregated
	Reporter[aggregated any]            func(aggregated, Translate, io.Writer) error

	storer interface {
		NewWriter(pref, ext string) (io.WriteCloser, error)
	}

	study[samples any, aggregated any] struct {
		Subjects map[string]Analyses

		collect   Collector[samples]
		aggregate Aggregator[samples, aggregated]
		report    Reporter[aggregated]
		store     storer
	}
)
