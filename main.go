package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
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

func main() {
	articles := parseArticles("./data/articles.csv")
	entities := parseEntities("./data/entities.csv")

	// TEST START
	articles = articles[0:2]
	// TEST END

	ctx := context.Background()
	assocArticles := processArticles(ctx, articles, entities)
	log.Println(assocArticles)
}

func processArticles(ctx context.Context, articles [][]string, entities [][]string) map[tokenize.Token]float64 {
	var (
		assocArticlesAccum = make(map[tokenize.Token][]float64)
		assocArticles      = make(map[tokenize.Token]float64)
	)
	// For Trump, Putin, Obama
	for _, entities := range entities {
		for _, article := range articles {
			eachText(article, func(text string) {
				switch text {
				case "":
					return
				}

				assocArticle, err := assocEntities(ctx, text, entities, tokenize.ANY)
				if err != nil {
					logAndFail(err)
				}
				for tok := range assocArticle {
					if dist, ok := assocArticlesAccum[tok]; ok {
						assocArticlesAccum[tok] = append(assocArticlesAccum[tok], dist...)
					}
				}
			})
		}
	}

	for tok, dist := range assocArticlesAccum {
		assocArticles[tok] = avgFloat(dist)
	}
	return assocArticles
}

func writeResult(path string, res map[tokenize.Token]float64) {
	file, err := os.Create(path)
	if err != nil {
		logAndFail(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()
	for token, dist := range res {
		record := []string{
			token.Text, fmt.Sprintf("%v", dist),
		}
		if err := w.Write(record); err != nil {
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

func assocEntities(ctx context.Context, text string, entities []string, pos tokenize.PoS) (assocEntities map[tokenize.Token]float64, err error) {
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
