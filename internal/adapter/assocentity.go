package internal

import (
	"context"
	"os"

	"github.com/ndabAP/assocentity/v9"
	"github.com/ndabAP/assocentity/v9/nlp"
	"github.com/ndabAP/assocentity/v9/tokenize"
)

func AssocEntities(ctx context.Context, text string, entities []string, pos tokenize.PoS) (assocEntities map[string]float64, err error) {
	// Create a NLP instance
	var credentialsFile = os.Getenv("GOOGLE_NLP_SERVICE_ACCOUNT_FILE_LOCATION")
	nlpTok := nlp.NewNLPTokenizer(credentialsFile, nlp.AutoLang)

	// Allow any part of speech
	posDeterm := nlp.NewNLPPoSDetermer(pos)

	// Do calculates the average distances
	assocEntities, err = assocentity.Do(ctx, nlpTok, posDeterm, text, entities)
	if err != nil {
		return
	}

	return
}
