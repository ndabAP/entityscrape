package api

import "context"

type APIResponse struct {
	Texts []string
}

func Fetch(ctx context.Context, entities []string) (res *APIResponse, err error) {
	res.Texts = make([]string, 0)
	for _, entity := range entities {
		// Search in API for entity
		var text string
		text, err = fetch(ctx, entity)
		if err != nil {
			return
		}

		res.Texts = append(res.Texts, text)
	}

	return
}

func fetch(ctx context.Context, entity string) (text string, err error) {
	return
}
