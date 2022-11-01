package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ndabAP/assocentity/v9"
	"github.com/ndabAP/assocentity/v9/nlp"
	"github.com/ndabAP/assocentity/v9/tokenize"
)

var logger = log.Default()

func init() {
	log.SetFlags(0)
}

var (
	gogSvcLocF = flag.String("gog-svc-loc", "", "")
)

func main() {
	flag.Parse()

	records, err := parseData("./data/news-mixed-2017.csv")
	if err != nil {
		logAndFail(err)
	}

	entities := []string{"Donald Trump", "Trump"}

	records = records[0:5]
	processData(records, entities)
}

func assocEntities(ctx context.Context, text string, entities []string, pos tokenize.PoS) (assocEntities map[string]float64, err error) {
	// Create a NLP instance
	nlpTok := nlp.NewNLPTokenizer(*gogSvcLocF, nlp.AutoLang)

	posDeterm := nlp.NewNLPPoSDetermer(pos)

	// Do calculates the average distances
	assocEntities, err = assocentity.Do(ctx, nlpTok, posDeterm, text, entities)
	if err != nil {
		return
	}
	return
}

func parseData(path string) (records [][]string, err error) {
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

func processData(records [][]string, entities []string) (err error) {
	each(records, func(content string) {
		text := content

		ctx := context.Background()
		pos := tokenize.NOUN | tokenize.ADJ | tokenize.VERB
		assocEnt, err := assocEntities(ctx, text, entities, pos)
		if err != nil {
			panic(err)
		}

		if err := writeResult(assocEnt); err != nil {
			panic(err)
		}
	})
	return
}

func writeResult(assocEnt map[string]float64) (err error) {
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	for token, dist := range assocEnt {
		record := []string{
			token, fmt.Sprintf("%v", dist),
		}
		if err := w.Write(record); err != nil {
			panic(err)
		}
	}

	return
}

func each(records [][]string, handler func(content string)) {
	for _, record := range records {
		for i, field := range record {
			switch i {
			// Index
			case 0,
				// Date
				1,
				// Link
				2,
				// Title
				3:
				continue

			// Content
			case 4:
				handler(field)
			}
		}
	}
}

func logAndFail(err error) {
	logger.Fatal(err)
}
