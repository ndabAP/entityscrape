package cli

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ndabAP/assocentity/v3"
	"github.com/ndabAP/assocentity/v3/tokenize"
)

const (
	sep = " "
)

var (
	credentialsFile string
)

func init() {
	godotenv.Load()

	credentialsFile = os.Getenv("GOOGLE_NLP_SERVICE_ACCOUNT_FILE_LOCATION")
}

func assoc(text string, entities []string) (map[string]float64, error) {
	// Create a NLP instance
	nlp, err := tokenize.NewNLP(credentialsFile, text, entities, false)
	if err != nil {
		return map[string]float64{}, err
	}

	// Join merges the entities with a simple algorithm
	dj := tokenize.NewDefaultJoin(sep)
	if err = dj.Join(nlp); err != nil {
		return map[string]float64{}, err
	}

	// Assoc calculates the average distances
	assocentities, err := assocentity.Assoc(dj, nlp, entities)
	if err != nil {
		return map[string]float64{}, err
	}

	return assocentities, err
}
