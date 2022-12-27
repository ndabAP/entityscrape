package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/ndabAP/assocentity/v10"
	"github.com/ndabAP/assocentity/v10/nlp"
	"github.com/ndabAP/assocentity/v10/tokenize"
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

type JSONRes struct {
	Els []JSONResEl // Marshal json map token
}
type JSONResEl struct {
	Entities         []string
	AssocEntitiesRes map[tokenize.Token]float64
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

	// For Trump, Putin, Obama
	for _, entities := range entities {
		assocEntitiesRes := make(map[tokenize.Token][]float64)

		var jsonResEl JSONResEl = JSONResEl{
			Entities:         entities,
			AssocEntitiesRes: make(map[tokenize.Token]float64),
		}

		for _, article := range articles {
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
					assocEntities, err := assocEntitiesDo(context.TODO(), text, entities, tokenize.ANY)
					if err != nil {
						logAndFail(err)
					}
					for tok := range assocEntities {
						if dist, ok := assocEntities[tok]; ok {
							assocEntitiesRes[tok] = append(assocEntitiesRes[tok], dist)
						}
					}
				}
			}
		}

		for tok, dist := range assocEntitiesRes {
			jsonResEl.AssocEntitiesRes[tok] = avgFloat(dist)
		}

		if len(jsonResEl.AssocEntitiesRes) == 0 {
			continue
		}

		file, err := json.MarshalIndent(jsonResEl, "", " ")
		if err != nil {
			logAndFail(err)
		}

		if err := ioutil.WriteFile(entities[0], file, 0644); err != nil {
			logAndFail(err)
		}
	}
}

// eachText iterates through articles and call textHandler func on every
// text containing column, which is in the current case at column 6
func eachText(article []string, textHandler func(content string)) {
	for idx, field := range article {
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
			textHandler(field)
		}
	}
}

func assocEntitiesDo(ctx context.Context, text string, entities []string, pos tokenize.PoS) (assocEntities map[tokenize.Token]float64, err error) {
	nlpTok := nlp.NewNLPTokenizer(*gogSvcLocF, nlp.AutoLang)
	posDeterm := nlp.NewNLPPoSDetermer(pos)
	assocEntities, err = assocentity.Do(ctx, nlpTok, posDeterm, text, entities)
	return
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
