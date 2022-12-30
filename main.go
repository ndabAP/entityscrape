// Donald_Trump.json
// {
// 	[pos]: [
// 		{
// 			word: string
// 			distance: number
// 		}
// 	]
// }

// Pre-format has duplicate words and accumulates

package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/ndabAP/assocentity/v11"
	"github.com/ndabAP/assocentity/v11/nlp"
	"github.com/ndabAP/assocentity/v11/tokenize"
)

func init() {
	log.SetFlags(0)
}

func init() {
	flag.Parse()
}

var (
	gogSvcLocF = flag.String("gog-svc-loc", "", "")
)

// TODO Doesn't take all articles into account. Last is overwritten. Need slice
type assocEntity struct {
	Entity       []string
	Associations map[tokenize.Token]float64
}

func (ea assocEntity) MarshalJSON() ([]byte, error) {
	var poSMapIds = map[tokenize.PoS]string{
		tokenize.UNKN:  "UNKNOWN",
		tokenize.ADJ:   "ADJ",
		tokenize.ADP:   "ADP",
		tokenize.ADV:   "ADV",
		tokenize.CONJ:  "CONJ",
		tokenize.DET:   "DET",
		tokenize.NOUN:  "NOUN",
		tokenize.NUM:   "NUM",
		tokenize.PRON:  "PRON",
		tokenize.PRT:   "PRT",
		tokenize.PUNCT: "PUNCT",
		tokenize.VERB:  "VERB",
		tokenize.X:     "X",
		tokenize.AFFIX: "AFFIX",
	}

	type assocEntityJSON struct {
		Entity       []string `json:"entity"`
		Associations map[string]struct {
			Distance     float64 `json:"distance"`
			PartOfSpeech string  `json:"partOfSpeech"`
		} `json:"associations"`
	}
	assocEntityRes := &assocEntityJSON{
		Entity: ea.Entity,
		Associations: make(map[string]struct {
			Distance     float64 `json:"distance"`
			PartOfSpeech string  `json:"partOfSpeech"`
		}),
	}

	for token, distance := range ea.Associations {
		pos := poSMapIds[token.PoS]
		assocEntityRes.Associations[token.Text] = struct {
			Distance     float64 "json:\"distance\""
			PartOfSpeech string  "json:\"partOfSpeech\""
		}{
			Distance:     distance,
			PartOfSpeech: pos,
		}
	}

	return json.Marshal(assocEntityRes)
}

func main() {
	articles, err := readCSV("./data/articles.csv")
	if err != nil {
		logAndFail(err)
	}
	// Remove CSV header
	articles = articles[1:]
	entities, err := readCSV("./data/entities.csv")
	if err != nil {
		logAndFail(err)
	}

	// TEST START
	articles = articles[0:2]
	// TEST END

	// For [[Donal Trump, Trump], [Putin], [Obama], ...]
	for _, entities := range entities {
		assocEntitiesAccum := make(map[tokenize.Token][]float64)

		var assocEntities assocEntity = assocEntity{
			Entity:       entities,
			Associations: make(map[tokenize.Token]float64),
		}

		// For [[ID, TITLE, TEXT], [ID2, TITLE, TEXT], ...]
		for _, article := range articles {
			// Or: text := article[5]
			for idx, text := range article {
				switch idx {
				case
					// article_id
					0,
					// publish_date
					1,
					// article_source_link
					2,
					// title
					3,
					// subtitle
					4:
					continue

				// Text
				case 5:
					nlpTok := nlp.NewNLPTokenizer(*gogSvcLocF, nlp.AutoLang)
					posDeterm := nlp.NewNLPPoSDetermer(tokenize.ANY)
					assocEntities, err := assocentity.Do(context.TODO(), nlpTok, posDeterm, text, entities)
					if err != nil {
						logAndFail(err)
					}

					for tok := range assocEntities {
						if dist, ok := assocEntities[tok]; ok {
							assocEntitiesAccum[tok] = append(assocEntitiesAccum[tok], dist)
						}
					}
				}
			}
		}

		for tok, dist := range assocEntitiesAccum {
			assocEntities.Associations[tok] = avgFloat(dist)
		}

		file, err := json.MarshalIndent(assocEntities, "", " ")
		if err != nil {
			logAndFail(err)
		}
		if err := os.WriteFile("./public/"+entities[0]+".json", file, 0644); err != nil {
			logAndFail(err)
		}

		// Next entity
	}
}

func readCSV(path string) (records [][]string, err error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err = csvReader.ReadAll()
	if err != nil {
		return
	}
	return
}

func logAndFail(err error) {
	log.Fatal(err)
}

// Returns the average of a float slice
func avgFloat(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}
