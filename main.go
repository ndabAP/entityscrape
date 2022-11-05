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

func init() {
	flag.Parse()
}

var (
	gogSvcLocF = flag.String("gog-svc-loc", "", "")
)

func main() {
	records := parseRecords("./data/records.csv")
	entities := parseEntities("./data/entities.csv")

	// TEST
	records = records[0:2]

	resChan := make(chan map[string]float64)
	ctx := context.Background()
	go processData(ctx, records, entities, resChan)

	for res := range resChan {
		writeResult(res)
	}
}

func parseRecords(path string) (data [][]string) {
	data, err := readCSV(path)
	if err != nil {
		logAndFail(err)
	}
	// Remove header
	data = data[1:]
	return
}

func parseEntities(path string) (entities [][]string) {
	entities, err := readCSV(path)
	if err != nil {
		logAndFail(err)
	}
	return
}

func processData(ctx context.Context, records [][]string, entities [][]string, resCh chan map[string]float64) {
	defer func() {
		close(resCh)
	}()
	for _, entities := range entities {
		eachText(records, func(text string) {
			pos := tokenize.ANY
			assocEnt, err := assocEntities(ctx, text, entities, pos)
			if err != nil {
				logAndFail(err)
			}

			resCh <- assocEnt
		})
	}
}

func writeResult(assocEnt map[string]float64) {
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	for token, dist := range assocEnt {
		record := []string{
			token, fmt.Sprintf("%v", dist),
		}
		if err := w.Write(record); err != nil {
			logAndFail(err)
		}
	}
}

func eachText(records [][]string, textHandler func(content string)) {
	for _, record := range records {
		for idx, field := range record {
			switch idx {
			// article_id
			case 0,
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
}

func assocEntities(ctx context.Context, text string, entities []string, pos tokenize.PoS) (assocEntities map[string]float64, err error) {
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
	logger.Fatal(err)
}
