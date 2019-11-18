package cli

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ndabAP/assocentity/v7"
	"github.com/ndabAP/assocentity/v7/tokenize"
)

const (
	lang = "en"
)

var (
	credentialsFile string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	credentialsFile = os.Getenv("GOOGLE_NLP_SERVICE_ACCOUNT_FILE_LOCATION")
}

// AssocEntities represents associated entities
type AssocEntities struct{}

// AssocEntities returns associated entities
func (ae AssocEntities) AssocEntities(text string, entities []string, logger *log.Logger) (map[tokenize.Token]float64, error) {
	// Create a NLP instance
	nlp, err := tokenize.NewNLP(credentialsFile, text, entities, lang)
	if err != nil {
		return map[tokenize.Token]float64{}, err
	}

	// Allow any part of speech
	psd := tokenize.NewPoSDetermer(tokenize.ANY)

	log.Printf("getting associations for aliases: %s", strings.Join(entities, ", "))

	assocEntities, err := assocentity.Do(nlp, psd, entities)
	if err != nil {
		log.Fatal(err)
	}

	return assocEntities, nil
}
