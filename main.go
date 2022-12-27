package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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
	articles := parseArticles("./data/articles.csv")
	entities := parseEntities("./data/entities.csv")

	// TEST START
	articles = articles[0:2]
	// TEST END

	assocEntitiesRes := make(map[tokenize.Token][]float64)
	jsonRes := JSONRes{
		Els: make([]JSONResEl, 0),
	}

	// For Trump, Putin, Obama
	for _, entities := range entities {
		var jsonResEl JSONResEl = JSONResEl{
			Entities:         entities,
			AssocEntitiesRes: make(map[tokenize.Token]float64),
		}

		for _, article := range articles {
			eachText(article, func(text string) {
				switch text {
				case "":
					return
				}

				assocEntities, err := assocEntitiesDo(context.TODO(), text, entities, tokenize.ANY)
				if err != nil {
					logAndFail(err)
				}
				for tok := range assocEntities {
					if dist, ok := assocEntitiesRes[tok]; ok {
						assocEntitiesRes[tok] = append(assocEntitiesRes[tok], dist...)
					}
				}
			})
		}

		for tok, dist := range assocEntitiesRes {
			jsonResEl.AssocEntitiesRes[tok] = avgFloat(dist)
		}

		// TODO Write here already JSON

		jsonRes.Els = append(jsonRes.Els, jsonResEl)
	}

	for i, el := range jsonRes.Els {
		file, _ := json.MarshalIndent(el, "", " ")

		_ = ioutil.WriteFile(strconv.FormatInt(int64(i), 10), file, 0644)
	}
}

func processArticles(ctx context.Context, articles [][]string, entities [][]string) JSONArrRes {
	assocEntitiesRes := make(map[tokenize.Token][]float64)
	jsonResArr := make([]JSONRes, 0)

	// For Trump, Putin, Obama
	for _, entities := range entities {
		var jsonRes JSONRes = JSONRes{
			Entities:         entities,
			AssocEntitiesRes: map[tokenize.Token]float64{},
		}

		for _, article := range articles {
			eachText(article, func(text string) {
				switch text {
				case "":
					return
				}

				assocEntities, err := assocEntitiesDo(ctx, text, entities, tokenize.ANY)
				if err != nil {
					logAndFail(err)
				}
				for tok := range assocEntities {
					if dist, ok := assocEntitiesRes[tok]; ok {
						assocEntitiesRes[tok] = append(assocEntitiesRes[tok], dist...)
					}
				}
			})
		}

		for tok, dist := range assocEntitiesRes {
			jsonRes.AssocEntitiesRes[tok] = avgFloat(dist)
		}

		jsonResArr = append(jsonResArr, jsonRes)
	}

	return jsonResArr
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

func parseArticles(path string) (articles [][]string) {
	articles, err := readCSV(path)
	if err != nil {
		logAndFail(err)
	}
	// Remove header
	articles = articles[1:]
	return
}

func parseEntities(path string) (entities [][]string) {
	entities, err := readCSV(path)
	if err != nil {
		logAndFail(err)
	}
	return
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
