package translator

import (
	"context"
	"fmt"
	"log/slog"
	"slices"

	translate "cloud.google.com/go/translate"
	"golang.org/x/sync/semaphore"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type (
	translator struct {
		client *translate.Client
		ctx    context.Context
	}
)

var (
	sema  = semaphore.NewWeighted(5)
	cache = make(map[[2]language.Tag]map[string]string)
)

func NewGoogle(ctx context.Context, creds string) translator {
	c, err := translate.NewClient(ctx, option.WithCredentialsFile(creds))
	if err != nil {
		panic(fmt.Errorf("gcloud translator client: %s", err))
	}
	return translator{
		client: c,
		ctx:    ctx,
	}
}

func ClearCache() {
	clear(cache)
}

func (translator translator) Translate(inputs []string, src, target language.Tag) ([]string, error) {
	slog.Debug("translate", "n", len(inputs), "src", src.String(), "target", target.String())

	sema.Acquire(translator.ctx, 1)
	defer sema.Release(1)

	ctx := translator.ctx

	outputs := make([]string, 0, len(inputs))

	// Check cache first.
	for i, input := range inputs {
		output, ok := cache[[2]language.Tag{src, target}][input]
		if !ok {
			continue
		}

		slog.Debug("cache hit", "input", input, "output", output)

		outputs = append(outputs, output)
		inputs = slices.Delete(inputs, i, i+1)
	}
	// Use cache exhaustively.
	if len(inputs) == 0 {
		return outputs, nil
	}

	opts := &translate.Options{
		Format: translate.Text,
		Source: src,
	}
	translation, err := translator.client.Translate(ctx, inputs, target, opts)
	if err != nil {
		return outputs, err
	}
	if len(translation) == 0 {
		return outputs, fmt.Errorf("translate: %s", inputs)
	}

	for i, t := range translation {
		output := t.Text
		outputs = append(outputs, output)

		// Add to cache.
		cache[[2]language.Tag{src, target}][inputs[i]] = output
	}
	return outputs, nil
}
