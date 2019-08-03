package cli

import (
	"log"
	"os"
	"strings"

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

// AssocEntities represents associated entities
type AssocEntities struct{}

// AssocEntities returns associated entities
func (ae AssocEntities) AssocEntities(text string, entities []string, logger *log.Logger) (map[string]float64, error) {
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

	log.Printf("getting associations for entities: %s", strings.Join(entities, ", "))

	// Assoc calculates the average distances
	assocEntities, err := assocentity.Assoc(dj, nlp, entities)
	if err != nil {
		return map[string]float64{}, err
	}

	return assocEntities, err
}
