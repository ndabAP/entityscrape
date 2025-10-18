// International sentiment of public figures
package isopf

import (
	"context"
	"encoding/json"
	"io"
	"slices"
	"sort"
	"strings"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/parser"
	"golang.org/x/text/language"
)

type (
	samples struct {
		heads, dependents []*tokenize.Token
	}
	aggregate struct {
		Word [2]string `json:"word"`
		N    int       `json:"n"`
	}
	aggregates struct {
		Heads      []aggregate `json:"heads"`
		Dependents []aggregate `json:"dependents"`
	}
)

var (
	ident = "isopf"

	collector = func(frames assocentity.Frames) samples {
		samples := samples{
			heads:      make([]*tokenize.Token, 0),
			dependents: make([]*tokenize.Token, 0),
		}
		f := func(tag tokenize.PartOfSpeechTag) bool {
			switch tag {
			case tokenize.PartOfSpeechTagAdj, tokenize.PartOfSpeechTagNoun, tokenize.PartOfSpeechTagVerb:
				return true
			default:
				return false
			}
		}
		frames.Forest().Heads(func(t *tokenize.Token) bool {
			if f(t.PartOfSpeech.Tag) {
				samples.heads = append(samples.heads, t)
			}
			return true
		})
		frames.Forest().Dependents(func(t *tokenize.Token) bool {
			if f(t.PartOfSpeech.Tag) {
				samples.dependents = append(samples.dependents, t)
			}
			return true
		})

		return samples
	}
	aggregator = func(s samples) aggregates {
		f := func(samples []*tokenize.Token) []aggregate {
			aggregates := make([]aggregate, 0)
			for _, sample := range samples {
				word := strings.ToLower(sample.Lemma)

				i := slices.IndexFunc(aggregates, func(aggr aggregate) bool {
					return aggr.Word[0] == word
				})
				// Find matches
				switch i {
				case -1:
					n := 1
					aggregates = append(aggregates, aggregate{
						Word: [2]string{word},
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
		return aggregates{
			Heads:      f(s.heads),
			Dependents: f(s.dependents),
		}
	}
	reporter = func(aggrs aggregates, translate cases.Translate, writer io.Writer) error {
		// Collect words to translate.
		f := func(aggregates []aggregate) error {
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
			return nil
		}
		if err := f(aggrs.Heads); err != nil {
			return err
		}
		if err := f(aggrs.Dependents); err != nil {
			return err
		}

		return json.NewEncoder(writer).Encode(&aggrs)
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
		parser = parser.AMND
		ext    = "json"

		filenames = make([]string, 0)
	)
	if err := cases.WalkCorpus("amnd", func(filename string) error {
		filenames = append(filenames, filename)
		return nil
	}); err != nil {
		return err
	}

	// Donald Trump
	// {
	// 	var (
	// 		ident  = "Trump"
	// 		entity = []string{ident, "Donald Trump", "Donald J. Trump", "Donald John Trump"}
	// 	)
	// 	study.Subjects[ident] = cases.Analyses{
	// 		Entity:        entity,
	// 		Ext:           ext,
	// 		Feats:         feats,
	// 		Filenames:     filenames,
	// 		FuzzyMatching: true,
	// 		Language:      lang,
	// 		Parser:        parser,
	// 	}
	// }
	// Elon Musk
	{
		var (
			ident  = "Musk"
			entity = []string{ident, "Elon Musk", "Elon Reeve Musk", "Elon R. Musk"}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:        entity,
			Ext:           ext,
			Feats:         feats,
			Filenames:     filenames,
			FuzzyMatching: true,
			Language:      lang,
			Parser:        parser,
		}
	}
	// Joe Biden
	{
		var (
			ident  = "Biden"
			entity = []string{ident, "Joe Biden", "Joseph Robinette Biden", "Joseph R. Biden", "Joseph Biden"}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:        entity,
			Ext:           ext,
			Feats:         feats,
			Filenames:     filenames,
			FuzzyMatching: true,
			Language:      lang,
			Parser:        parser,
		}
	}
	// Vladimir Putin
	{
		var (
			ident  = "Putin"
			entity = []string{ident, "Vladimir Putin", "Vladimir Vladimirovich Putin"}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:        entity,
			Ext:           ext,
			Feats:         feats,
			Filenames:     filenames,
			FuzzyMatching: true,
			Language:      lang,
			Parser:        parser,
		}
	}

	if err := study.Conduct(ctx); err != nil {
		return err
	}

	return nil
}
