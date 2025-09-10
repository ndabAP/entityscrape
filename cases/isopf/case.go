// International sentiment of public figures
package isopf

import (
	"context"
	"encoding/json"
	"io"
	"slices"
	"sort"
	"unicode"
	"unicode/utf8"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/entityscrape/cases"
	"github.com/ndabAP/entityscrape/parser"
	"golang.org/x/text/language"
)

type (
	samples struct {
		ancestors, descendants []*tokenize.Token
	}
	aggregate struct {
		Word [2]string `json:"heads"`
		N    int       `json:"n"`
	}
	aggregates struct {
		Ancestors   []aggregate `json:"ancestors"`
		Descendants []aggregate `json:"descendants"`
	}
)

var (
	ident = "isopf"

	collector = func(analyses assocentity.Analyses) samples {
		var (
			ancestors   = analyses.Forest().Ancestors(nil)
			descendants = analyses.Forest().Descendants(nil)
		)
		// Reduce
		del := func(token *tokenize.Token) bool {
			switch token.PartOfSpeech.Tag {
			case tokenize.PartOfSpeechTagAdj, tokenize.PartOfSpeechTagNoun, tokenize.PartOfSpeechTagVerb:
			default:
				return true
			}

			// Ignore non-ASCII characters.
			r, _ := utf8.DecodeRuneInString(token.Lemma)
			switch {
			case unicode.IsDigit(r), unicode.IsLetter(r):
			default:
				return true
			}

			return false
		}
		return samples{
			ancestors:   slices.DeleteFunc(ancestors, del),
			descendants: slices.DeleteFunc(descendants, del),
		}
	}
	aggregator = func(s samples) aggregates {
		f := func(samples []*tokenize.Token) []aggregate {
			aggregates := make([]aggregate, 0)
			for _, sample := range samples {
				i := slices.IndexFunc(aggregates, func(aggr aggregate) bool {
					return aggr.Word[0] == sample.Lemma
				})
				// Find matches
				switch i {
				case -1:
					n := 1
					aggregates = append(aggregates, aggregate{
						Word: [2]string{sample.Lemma},
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
			Ancestors:   f(s.ancestors),
			Descendants: f(s.descendants),
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
		f(aggrs.Ancestors)
		f(aggrs.Descendants)

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
	// Elon Musk
	{
		var (
			ident  = "Musk"
			entity = []string{ident, "Elon Musk", "Elon Reeve Musk"}
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
	// Joe Biden
	{
		var (
			ident  = "Biden"
			entity = []string{ident, "Joe Biden", "Joseph Robinette Biden", "Joseph R. Biden", "Joseph Biden"}
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
	// Vladimir Putin
	{
		var (
			ident  = "Putin"
			entity = []string{ident, "Vladimir Putin", "Vladimir Vladimirovich Putin"}
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

	if err := study.Conduct(ctx); err != nil {
		return err
	}

	return nil
}
