package translator

import (
	"context"
	"fmt"

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

var sema = semaphore.NewWeighted(5)

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

func (t translator) Translate(inputs []string, src, target language.Tag) ([]string, error) {
	sema.Acquire(t.ctx, 1)
	defer sema.Release(1)

	var (
		ctx  = t.ctx
		opts = &translate.Options{
			Format: translate.Text,
			Source: src,
		}

		outputs = make([]string, 0, len(inputs))
	)
	translation, err := t.client.Translate(ctx, inputs, target, opts)
	if err != nil {
		return outputs, err
	}
	if len(translation) == 0 {
		return outputs, fmt.Errorf("translate: %s", inputs)
	}

	for _, t := range translation {
		outputs = append(outputs, t.Text)
	}
	return outputs, nil
}
