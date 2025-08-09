// International sentiment of public figures
package isopf

import (
	"context"
	"encoding/json"
	"io"
	"path/filepath"
	"slices"
	"sort"
	"unicode"
	"unicode/utf8"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/dependency"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/parser"
	"golang.org/x/text/language"
)

const depth = 3

type (
	sample    [depth]*tokenize.Token
	aggregate struct {
		Heads [depth][2]string `json:"heads"`
		N     int              `json:"n"`
	}
	aggregates []aggregate
)

var (
	ident = "isopf"

	collector = func(analyses assocentity.Analyses) []sample {
		var (
			entities = analyses.Forest().Entities()
			samples  = make([]sample, 0)
		)

		analyses.Forest().Walk(func(token *tokenize.Token, tree dependency.Tree) bool {
			if slices.Contains(entities, token) {
				return true
			}

			var (
				s sample
				d int = 1
			)
			tree.(token, func(t *tokenize.Token) bool {
				if d == depth {
					return false
				}

				// Ignore multi-token entity.
				if slices.Contains(entities, token) {
					return true
				}
				// Ignore possesive noun suffix.
				switch token.Lemma {
				case "’s", "'s", "s":
					return true
				default:
				}
				// Ignore non-ASCII characters.
				r, _ := utf8.DecodeRuneInString(token.Lemma)
				switch {
				case unicode.IsDigit(r), unicode.IsLetter(r):
				default:
					return true
				}

				s[d-1] = token
				d++

				return true
			})

			samples = append(samples, s)

			return false
		})

		// Delete samples that don't contain all depth tokens.
		samples = slices.DeleteFunc(samples, func(sample sample) bool {
			for _, t := range sample {
				if t == nil {
					return true
				}
			}

			return false
		})
		return samples
	}
	aggregator = func(samples []sample) aggregates {
		aggregates := make(aggregates, 0, len(samples))
		for _, sample := range samples {
			ws := make([]string, depth)
			for i, w := range sample {
				ws[i] = w.Lemma
			}

			i := slices.IndexFunc(aggregates, func(aggregate aggregate) bool {
				for i, w := range aggregate.Heads {
					if w[0] != ws[i] {
						return false
					}
				}

				return true
			})
			// Find matches
			switch i {
			case -1:
				n := 1
				aggregates = append(aggregates, aggregate{
					Heads: [depth][2]string{
						{ws[0], ""},
						{ws[1], ""},
						{ws[2], ""},
					},
					N: n,
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
		for i, aggregate := range aggregates {
			words = append(words, aggregate.Heads[i][0])
		}
		w, err := translate(words)
		if err != nil {
			return err
		}
		// Add translated words back.
		for i := range aggregates {
			aggregates[i].Heads[i][1] = w[i]
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
		parser = parser.AMND
		ext    = "json"

		filenames = make([]string, 0)
	)
	if err := cases.WalkCorpus("amnd", func(filename string) error {
		if filepath.Ext(filename) != ".json" {
			return nil
		}

		filenames = append(filenames, filename)
		return nil
	}); err != nil {
		return err
	}

	// Donald Trump
	{
		var (
			ident  = "Trump"
			entity = []string{ident, "Donald Trump", "Donald J. Trump", "Donald John Trump"}
		)
		study.Subjects[ident] = cases.Analyses{
			Entity:    entity,
			Ext:       ext,
			Feats:     feats,
			Filenames: filenames,
			Reduct:    true,
			Language:  lang,
			Parser:    parser,
		}

	}
	// // Google
	// {
	// 	var (
	// 		ident  = "Google"
	// 		entity = []string{ident}
	// 	)
	// 	study.Subjects[ident] = cases.Analyses{
	// 		Entity:    entity,
	// 		Ext:       ext,
	// 		Feats:     feats,
	// 		Filenames: filenames,
	// 		Language:  lang,
	// 		Parser:    parser,
	// 	}

	// }
	// // Amazon
	// {
	// 	var (
	// 		ident  = "Amazon"
	// 		entity = []string{ident}
	// 	)
	// 	study.Subjects[ident] = cases.Analyses{
	// 		Entity:    entity,
	// 		Ext:       ext,
	// 		Feats:     feats,
	// 		Filenames: filenames,
	// 		Language:  lang,
	// 		Parser:    parser,
	// 	}
	// }

	if err := study.Conduct(ctx); err != nil {
		return err
	}

	return nil
}
