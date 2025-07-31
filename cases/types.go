package cases

import (
	"io"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"golang.org/x/text/language"
)

type (
	Parser   func(io.Reader) ([]string, error)
	Analyses struct {
		Entity    []string
		Feats     tokenize.Features
		Filenames []string
		Parser    Parser
		Language  language.Tag
		Ext       string
	}

	Translate func([]string) ([]string, error)

	Collector[samples any]              func(assocentity.Analyses) samples
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
