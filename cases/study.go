package cases

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"path"
	"path/filepath"

	"github.com/ndabAP/assocentity"
	"github.com/ndabAP/assocentity/tokenize"
	"github.com/ndabAP/assocentity/tokenize/nlp"
	"github.com/ndabAP/entityscrape/parser"
	"github.com/ndabAP/entityscrape/store"
	"github.com/ndabAP/entityscrape/translator"
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

	Collector[samples any]              func(assocentity.Analyses) samples
	Aggregator[samples, aggregated any] func(samples) aggregated
	Reporter[aggregated any]            func(aggregated, Translate, io.Writer) error

	Translate func([]string) ([]string, error)
	storer    interface {
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

var (
	Corpus                   fs.FS
	GoogleCloudSvcAccountKey string
)

func NewStudy[samples, aggregated any](
	ident string,
	collect Collector[samples],
	aggregate Aggregator[samples, aggregated],
	report Reporter[aggregated],
) study[samples, aggregated] {
	store := store.NewFile(filepath.Join("cases", ident, "report"))
	subjects := make(map[string]Analyses)
	return study[samples, aggregated]{
		Subjects:  subjects,
		collect:   collect,
		aggregate: aggregate,
		report:    report,
		store:     store,
	}
}

func (study study[samples, aggregated]) Conduct(ctx context.Context) error {
	translator := translator.NewGoogle(ctx, GoogleCloudSvcAccountKey)

	slog.Debug("processing subjects", "n", len(study.Subjects))
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
		slog.Debug("reporting done")
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

		slog.Debug("processing file", "filename", filename)
		file, err := Corpus.Open(path.Join("corpus", filename))
		if err != nil {
			return assocentity.Analyses{}, err
		}
		text, err := parse(file)
		if errors.Is(err, parser.ErrTextTooShort) {
			slog.Debug("skipping short text")
			continue
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

func WalkCorpus(corpus string, fn func(filename string) error) {
	root := filepath.Join("corpus", corpus)
	err := fs.WalkDir(Corpus, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		return fn(path)
	})
	if err != nil {
		panic(err)
	}
}
