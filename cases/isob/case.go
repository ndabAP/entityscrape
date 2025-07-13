// International sentiment of brands
package isob

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"
	"slices"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/dependency"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/parser"
	"golang.org/x/text/language"
)

var logger = slog.Default()

const lvl = 3 // Indices are 0-based

type (
	sample    [lvl]*tokenize.Token
	aggregate struct {
		Heads [lvl][2]string `json:"heads"`
		N     int            `json:"n"`
	}
	aggregates []aggregate
)

var (
	ident = "isob"

	collector = func(analyses assocentity.Analyses) []sample {
		var (
			entities = analyses.Forest().Entities()
			samples  = make([]sample, 0)
		)

		analyses.Forest().Walk(func(token *tokenize.Token, tree dependency.Tree) bool {
			if !slices.Contains(entities, token) {
				return true
			}

			var (
				s sample
				l int
			)
			tree.Ancestors(token, func(token *tokenize.Token) bool {
				if l == lvl {
					return false
				}

				s[l-1] = token

				l++
				return true
			})
			samples = append(samples, s)

			return false
		})

		return samples
	}
	aggregator = func(samples []sample) aggregates {
		aggregates := make(aggregates, 0, len(samples))
		for _, sample := range samples {
			ws := [lvl]string{}
			for i, w := range sample {
				ws[i] = w.Lemma
			}

			i := slices.IndexFunc(aggregates, func(aggregate aggregate) bool {
				return ws == aggregate.Heads[0]
			})
			// Find matches
			switch i {
			case -1:
				var (
					heads = ws
					n     = 1
				)
				aggregates = append(aggregates, aggregate{
					Heads: heads,
					N:     n,
				})
			// Found
			default:
				aggregates[i].N++
			}
		}

		return aggregates
	}
	reporter = func(aggregates aggregates, translate cases.Translate, writer io.Writer) error {
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

	feats := tokenize.FeatureSyntax

	var (
		lang      = language.English
		filenames = make([]string, 0)
		parser    = parser.AMND
		ext       = "json"
	)
	cases.WalkCorpus("amnd", func(filename string) error {
		if filepath.Ext(filename) != ".json" {
			return nil
		}

		filenames = append(filenames, filename)
		return nil
	})

	// Apple
	{
		var (
			ident  = "Apple"
			entity = []string{ident}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:    entity,
			Feats:     feats,
			Filenames: filenames,
			Language:  lang,
			Parser:    parser,
			Ext:       ext,
		}

	}
	// Google
	{
		var (
			ident  = "Google"
			entity = []string{ident}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:    entity,
			Feats:     feats,
			Filenames: filenames,
			Language:  lang,
			Parser:    parser,
			Ext:       ext,
		}

	}
	// Amazon
	{
		var (
			ident  = "Amazon"
			entity = []string{ident}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:    entity,
			Feats:     feats,
			Filenames: filenames,
			Language:  lang,
			Parser:    parser,
			Ext:       ext,
		}
	}

	if err := study.Conduct(ctx); err != nil {
		return err
	}

	return nil
}
