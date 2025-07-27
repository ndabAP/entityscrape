package cases

import (
	"context"
	"errors"
	"log/slog"
	"math/rand/v2"
	"path"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/assocentity/tokenize/nlp"
	"github.com/ndabAP/entityscrape/parser"
	"github.com/ndabAP/entityscrape/translator"
	"golang.org/x/text/language"
)

func (study study[samples, aggregated]) Conduct(ctx context.Context) error {
	defer translator.ClearCache()

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
			feats     = analyses.Feats
			filenames = analyses.Filenames
			lang      = analyses.Language
			parser    = analyses.Parser
			ext       = analyses.Ext
		)
		tokenizer := nlp.New(GoogleCloudSvcAccountKey, lang.String())
		analyses, err := study.analysis(
			ctx,
			entity,
			filenames,
			parser,
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
				slog.Debug("skipping translation")
				return w, nil
			default:
			}

			return translator.Translate(w, lang, language.English)
		}
		writer, err := study.store.NewWriter(subject, ext)
		if err != nil {
			return err
		}
		if err := study.report(aggregated, translator, writer); err != nil {
			return err
		}
		writer.Close()

		slog.Debug("reporting done", "subject", subject)
	}

	return nil
}

func (study study[samples, aggregated]) analysis(
	ctx context.Context,
	entity,
	filenames []string,
	parse Parser,
	tokenize tokenize.Tokenizer,
	feats tokenize.Features,
) (
	assocentity.Analyses,
	error,
) {
	slog.Debug("parsing files", "n", len(filenames))
	texts := make([]string, 0, len(filenames))
	for _, filename := range filenames {
		select {
		case <-ctx.Done():
			return assocentity.Analyses{}, ctx.Err()
		default:
		}

		// Sampling
		n := rand.Uint64N(100)
		if n >= Sampling {
			slog.Debug("skipping file (sampling)", "filename", filename, "n", n)
			continue
		}

		slog.Debug("processing file", "filename", filename)
		file, err := Corpus.Open(path.Join("corpus", filename))
		if err != nil {
			return assocentity.Analyses{}, err
		}
		text, err := parse(file)
		switch {
		case errors.Is(err, parser.ErrTextTooShort):
			slog.Debug("skipping short text")
			continue
		case errors.Is(err, parser.ErrUnsupportedLang):
			slog.Debug("skipping unsupported language")
			continue
		default:
		}
		if err != nil {
			return assocentity.Analyses{}, err
		}
		texts = append(texts, text...)
	}
	slog.Debug("files parsed")

	slog.Debug("creating analyses")
	src := assocentity.NewSource(entity, texts)
	return src.Analyses(ctx, tokenize, feats)
}
