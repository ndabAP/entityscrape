// National sentiment of political speeches
package nsops

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"path/filepath"
	"slices"
	"sort"

	"cloud.google.com/go/language/apiv1/languagepb"
	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/dependency"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/parser"
	"golang.org/x/text/language"
)

type (
	sample struct {
		*tokenize.Token
		from bool
	}
	aggregate struct {
		Word [2]string `json:"word"`
		PoS  string    `json:"pos"`
		N    int       `json:"n"`
	}
	aggregates []aggregate
)

var (
	ident = "nsops"

	collector = func(analyses assocentity.Analyses) []sample {
		var (
			entities = analyses.Forest().Entities()
			samples  = make([]sample, 0)
		)
		walker := func(
			from,
			to *tokenize.Token,
			_ tokenize.DependencyEdgeLabel,
			tree dependency.Tree,
		) bool {
			switch {
			// To entity
			case slices.Contains(entities, to):
				switch to.PartOfSpeech.Tag {
				case tokenize.PartOfSpeechTagVerb, tokenize.PartOfSpeechTagNoun, tokenize.PartOfSpeechTagAdj:
					slog.Debug("adding sample", "word", to.Text.Content)
					samples = append(samples, sample{
						Token: to,
					})
				default:
					// Skip
				}

				return true

			// From entity
			case slices.Contains(entities, from):
				switch from.PartOfSpeech.Tag {
				case tokenize.PartOfSpeechTagVerb, tokenize.PartOfSpeechTagNoun, tokenize.PartOfSpeechTagAdj:
					slog.Debug("adding sample", "word", from.Text.Content)
					samples = append(samples, sample{
						Token: from,
						from:  true,
					})
				default:
				}
			}

			return true
		}
		analyses.Forest().Dependencies(walker)

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
					pos  = languagepb.PartOfSpeech_Tag_name[int32(sample.PartOfSpeech.Tag)]
					n    = 1
				)
				aggregates = append(aggregates, aggregate{
					Word: word,
					PoS:  pos,
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

	feats := tokenize.FeatureSyntax

	// Germany
	{
		lang := language.German
		entity := []string{"Deutschland", "Deutschlands", "Deutschlande"}

		// GPSC
		{
			var (
				filenames = []string{
					filepath.Join(cases.GetCorpusRootDir(), "gpsc", "Bundesregierung.xml"),
				}
				parser = parser.GPSC
			)
			study.Subjects["Deutschland"] = cases.Analyses{
				Entity:    entity,
				Feats:     feats,
				Filenames: filenames,
				Language:  lang,
				Parser:    parser,
				Reduct:    true,
				Ext:       "json",
			}
		}
	}
	// // Russia
	// {
	// 	lang := language.Russian
	// 	entity := []string{"Россия", "России", "Россией"}

	// 	// DS
	// 	{
	// 		parser := parser.DS
	// 		filenames := make([]string, 0)
	// 		if err := cases.WalkCorpus("ds", func(filename string) error {
	// 			filenames = append(filenames, filename)
	// 			return nil
	// 		}); err != nil {
	// 			return err
	// 		}
	// 		study.Subjects["Russia"] = cases.Analyses{
	// 			Entity:    entity,
	// 			Feats:     feats,
	// 			Filenames: filenames,
	// 			Language:  lang,
	// 			Parser:    parser,
	// 			Reduct:    true,
	// 			Ext:       "json",
	// 		}
	// 	}
	// }
	// // USA
	// {
	// 	lang := language.English
	// 	entity := []string{"United States", "USA", "United States of America"}

	// 	// CRFTC
	// 	{
	// 		parser := parser.CRFTC
	// 		filenames := make([]string, 0)
	// 		if err := cases.WalkCorpus("crftc", func(filename string) error {
	// 			filenames = append(filenames, filename)
	// 			return nil
	// 		}); err != nil {
	// 			return err
	// 		}
	// 		study.Subjects["USA"] = cases.Analyses{
	// 			Entity:    entity,
	// 			Feats:     feats,
	// 			Filenames: filenames,
	// 			Language:  lang,
	// 			Parser:    parser,
	// 			Reduct:    true,
	// 			Ext:       "json",
	// 		}
	// 	}
	// }

	if err := study.Conduct(ctx); err != nil {
		return err
	}

	return nil
}
