package cases

import (
	"bufio"
	"context"
	"errors"
	"log/slog"
	"math/rand/v2"
	"os"
	"strings"
	"unicode"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/assocentity/tokenize/nlp"
	"github.com/ndabAP/entityscrape/translator"
	"golang.org/x/text/language"
)

func (study study[samples, aggregated]) Conduct(ctx context.Context) error {
	slog.Debug("processing subjects", "n", len(study.Subjects))

	translator := translator.NewGoogle(ctx, GoogleCloudSvcAccountKey)
	for subject, analyses := range study.Subjects {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		slog.Debug("processing analyses", "subject", subject)
		var (
			entity    = analyses.Entity
			ext       = analyses.Ext
			feats     = analyses.Feats
			filenames = analyses.Filenames
			reduct    = analyses.Reduct
			lang      = analyses.Language
			parser    = analyses.Parser
		)
		tokenizer := nlp.New(GoogleCloudSvcAccountKey, lang.String())
		analyses, err := study.analysis(
			ctx,
			entity,
			filenames,
			parser,
			reduct,
			tokenizer,
			feats,
		)
		if err != nil {
			return err
		}
		slog.Debug("analysis done")

		slog.Debug("collecting samples")
		samples := study.collect(analyses)
		slog.Debug("sample collection done")

		slog.Debug("aggregating samples")
		aggregated := study.aggregate(samples)
		slog.Debug("aggregation done")

		slog.Debug("reporting aggregation")
		translator := func(w []string) ([]string, error) {
			switch lang {
			case language.English:
				slog.Debug("skipping translation for English")
				return w, nil
			default:
			}

			return translator.Translate(w, lang, language.English)
		}
		if err := func() error {
			pref := strings.Map(func(r rune) rune {
				switch {
				case r >= 'a' && r <= 'z':
					return r
				case r >= 'A' && r <= 'Z':
					return unicode.ToLower(r)
				case r == ' ', r == '-':
					return '_'
				default:
					return -1
				}
			}, subject)
			writer, err := study.store.NewWriter(pref, ext)
			if err != nil {
				return err
			}
			//nolint:errcheck
			defer writer.Close()

			if err := study.report(aggregated, translator, writer); err != nil {
				return err
			}

			return nil
		}(); err != nil {
			return err
		}
		slog.Debug("reporting done", "subject", subject)
	}

	return nil
}

func (study study[samples, aggregated]) analysis(
	ctx context.Context,
	entity,
	filenames []string,
	parser Parser,
	reduct bool,
	tokenizer tokenize.Tokenizer,
	feats tokenize.Features,
) (
	assocentity.Analyses,
	error,
) {
	slog.Debug("parsing files", "n", len(filenames))

	var (
		texts = make([]string, 0, len(filenames))

		textChan = make(chan []byte, 50)
		errChan  = make(chan error, 1)
	)

	// Consumer
	go func() {
		defer close(errChan)

		for text := range textChan {
			n := rand.Uint64N(100)
			if n >= SampleRate {
				continue
			}

			var err error
			if reduct {
				text, err = study.reduct(text, entity)
				if errors.Is(err, errEntityNotFound) {
					continue
				}
			}
			if err != nil {
				errChan <- err
				return
			}
			texts = append(texts, string(text))
		}
	}()
	// Producer
	go func() {
		defer close(textChan)

		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				errChan <- err
				return
			}
			for err := range parser(file, textChan) {
				if errors.Is(err, bufio.ErrTooLong) {
					continue
				}
				errChan <- err
			}
			//nolint:errcheck
			_ = file.Close()
		}
	}()

	select {
	case <-ctx.Done():
		return assocentity.Analyses{}, ctx.Err()

	case err := <-errChan:
		if err != nil {
			return assocentity.Analyses{}, err
		}
	}
	slog.Debug("texts sampled and parsed", "n", len(texts))

	slog.Debug("creating analyses")
	src := assocentity.NewSource(entity, texts)
	return src.Analyses(ctx, tokenizer, feats, assocentity.NFKC)
}
