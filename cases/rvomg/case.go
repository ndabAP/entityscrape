// Root verbs of music genres
package rvomg

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"slices"
	"sort"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/parser"
	"golang.org/x/text/language"
)

var logger = slog.Default()

type (
	sample    *tokenize.Token
	aggregate struct {
		Word [2]string `json:"word"`
		N    int       `json:"n"`
	}
	aggregates []aggregate
)

var (
	ident = "rvomg"

	collector = func(analyses assocentity.Analyses) []sample {
		roots := analyses.Forest().Roots()
		samples := make([]sample, 0, len(roots))
		for _, root := range roots {
			samples = append(samples, root)
		}

		return samples
	}
	aggregator = func(samples []sample) aggregates {
		aggregates := make(aggregates, 0, len(samples))
		for _, sample := range samples {
			w := sample.Lemma
			i := slices.IndexFunc(aggregates, func(aggregate aggregate) bool {
				return w == aggregate.Word[0]
			})
			// Find matches
			switch i {
			case -1:
				var (
					word = [2]string{w}
					n    = 1
				)
				aggregates = append(aggregates, aggregate{
					Word: word,
					N:    n,
				})
			// Found
			default:
				aggregates[i].N++
			}
		}

		// Top n sorted
		const limit = 10
		sort.Slice(aggregates, func(i, j int) bool {
			return aggregates[i].N > aggregates[j].N
		})
		if len(aggregates) > limit {
			aggregates = aggregates[:limit]
		}

		return aggregates
	}
	reporter = func(aggregates aggregates, translate cases.Translate, writer io.Writer) error {
		// Collect words to translate.
		words := make([]string, 0, len(aggregates))
		for _, aggregate := range aggregates {
			words = append(words, aggregate.Word[0])
		}
		w, err := translate(words)
		if err != nil {
			return err
		}
		// Add translated words back.
		for i := range aggregates {
			aggregates[i].Word[1] = w[i]
		}

		return json.NewEncoder(writer).Encode(&aggregates)
	}
)

func Conduct(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return conduct(ctx)
}

func conduct(ctx context.Context) error {
	study := cases.NewStudy(ident, collector, aggregator, reporter)

	var (
		feats  = tokenize.FeatureSyntax
		lang   = language.English
		parser = parser.Etc
	)

	// Pop
	{
		entity := []string{"Pop"}

		filenames := make([]string, 0)
		cases.WalkCorpus("etc/pop", func(filename string) error {
			filenames = append(filenames, filename)
			return nil
		})
		study.Subjects["Pop"] = cases.Analyses{
			Entity:    entity,
			Feats:     feats,
			Filenames: filenames,
			Language:  lang,
			Parser:    parser,
			Ext:       "json",
		}
	}

	if err := study.Conduct(ctx); err != nil {
		return err
	}

	return nil
}
