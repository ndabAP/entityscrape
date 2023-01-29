package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"

	"github.com/ndabAP/assocentity/v12"
	"github.com/ndabAP/assocentity/v12/nlp"
	"github.com/ndabAP/assocentity/v12/tokenize"
)

func init() {
	log.SetFlags(0)
	flag.Parse()
}

var (
	gogSvcLocF = flag.String(
		"gog-svc-loc",
		"",
		"Google Clouds NLP JSON service account file, example: -gog-svc-loc=\"~/gog-svc-loc.json\"",
	)
)

func main() {
	articles, err := readCSV("./source/articles.csv")
	if err != nil {
		log.Fatal(err)
	}
	entities, err := readCSV("./source/entities.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Validate
	if len(*gogSvcLocF) == 0 {
		log.Fatal("missing service account file")
	}
	if len(entities) == 0 {
		log.Fatal("missing entities")
	}

	// Accumulate texts
	texts := make([]string, 0)
	// Remove CSV header
	articles = articles[1:]
	// For [[ID1, TITLE1, TEXT1], [ID2, TITLE2, TEXT2], ...]
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
				texts = append(texts, text)
			}
		}
	}

	// TEST START
	texts = texts[0:20]
	// TEST END

	if len(texts) == 0 {
		log.Fatal("missing texts")
	}

	log.Printf("len(texts)=%d", len(texts))
	log.Printf("len(entities)=%d", len(entities))

	// Get mean distance per entity
	log.Println("get meanN")
	nlpTok := nlp.NewNLPTokenizer(*gogSvcLocF, nlp.AutoLang)
	for _, entities := range entities {
		log.Printf("entities=%v", entities)

		// First entity is primary one
		entity := entities[0]
		log.Printf("entity=%s", entity)

		meanN, err := assocentity.MeanN(
			context.Background(),
			nlpTok,
			tokenize.ANY,
			texts,
			entities,
		)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("len(meanN)=%d", len(meanN))

		// Convert to slice to make it sortable
		log.Println("convert to slice")
		type meanNVal struct {
			dist float64
			tok  tokenize.Token
		}
		meanNVals := make([]meanNVal, 0)
		for tok, dist := range meanN {
			// TODO Only allow a-zA-Z0-9 runes

			// Skip unknown pos
			switch tok.PoS {
			case tokenize.X, tokenize.UNKN:
				continue
			}

			meanNVals = append(meanNVals, meanNVal{
				dist: dist,
				tok:  tok,
			})
		}

		// Sort by closest distance
		log.Println("sort by pos and dist")
		sort.Slice(meanNVals, func(i, j int) bool {
			if meanNVals[i].tok.PoS != meanNVals[j].tok.PoS {
				return meanNVals[i].tok.PoS < meanNVals[j].tok.PoS
			}
			return meanNVals[i].dist < meanNVals[j].dist
		})

		// Top 10 per pos
		log.Println("limit top 10")
		type topMeanNVal struct {
			Dist float64 `json:"dist"`
			Pos  string  `json:"pos"`
			Text string  `json:"text"`
		}
		topMeanNVals := make([]topMeanNVal, 0)
		poSHits := make(map[tokenize.PoS]int)
		for _, meanNVal := range meanNVals {
			// Stop at 10 results
			if poSHits[meanNVal.tok.PoS] >= 10 {
				continue
			}

			topMeanNVals = append(topMeanNVals, topMeanNVal{
				Dist: meanNVal.dist,
				Pos:  tokenize.PoSMapStr[meanNVal.tok.PoS],
				Text: meanNVal.tok.Text,
			})

			poSHits[meanNVal.tok.PoS] += 1
		}
		log.Printf("len(topMeanNVals)=%d", len(topMeanNVals))

		// Write top 10 to disk
		log.Println("write to disk")
		file, err := json.MarshalIndent(&topMeanNVals, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		name := filepath.Join("public", url.QueryEscape(entity)+".json")
		if err := os.WriteFile(name, file, 0644); err != nil {
			log.Fatal(err)
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
