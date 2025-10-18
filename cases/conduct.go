package cases

import (
	"bufio"
	"bytes"
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
	"github.com/ndabAP/entityscrape/sbd"
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
			entity        = analyses.Entity
			ext           = analyses.Ext
			feats         = analyses.Feats
			filenames     = analyses.Filenames
			fuzzyMatching = analyses.FuzzyMatching
			lang          = analyses.Language
			parser        = analyses.Parser
		)
		tokenizer := nlp.New(GoogleCloudSvcAccountKey, lang.String())
		frames, err := study.frames(
			ctx,
			entity,
			filenames,
			parser,
			fuzzyMatching,
			tokenizer,
			feats,
			lang,
		)
		if err != nil {
			return err
		}
		slog.Debug("analysis done")

		slog.Debug("collecting samples")
		samples := study.collect(frames)
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

func (study study[samples, aggregated]) frames(
	ctx context.Context,
	entity,
	filenames []string,
	parser Parser,
	fuzzyMatching bool,
	tokenizer tokenize.Tokenizer,
	feats tokenize.Features,
	lang language.Tag,
) (
	assocentity.Frames,
	error,
) {
	slog.Debug("parsing files", "n", len(filenames))
	if fuzzyMatching {
		slog.Debug("fuzzy matching enabled")
	}

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

			t := string(text)

			if fuzzyMatching {
				var (
					buf = new(bytes.Buffer)

					c    = make(chan string, 50)
					done = make(chan struct{}, 1)
				)

				// Consumer
				study.fuzzyMatch(c, entity, buf, done)
				// Producer
				sbd.Tokenize(lang, t, c)
				close(c)

				<-done
				if buf.Len() > 0 {
					texts = append(texts, strings.TrimSpace(buf.String()))
				}

				continue
			}

			// No fuzzy matching.
			texts = append(texts, strings.TrimSpace(t))
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
		return assocentity.Frames{}, ctx.Err()

	case err := <-errChan:
		if err != nil {
			return assocentity.Frames{}, err
		}
	}
	slog.Debug("texts sampled and parsed", "n", len(texts))

	slog.Debug("generating frames")
	src := assocentity.NewSource(entity, texts)
	return src.Frames(ctx, tokenizer, feats, assocentity.NFKC)
}
