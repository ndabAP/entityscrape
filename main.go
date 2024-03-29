package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/ndabAP/assocentity/v14"
	"github.com/ndabAP/assocentity/v14/nlp"
	"github.com/ndabAP/assocentity/v14/tokenize"
)

func init() {
	log.SetFlags(0)
	flag.Parse()
}

var (
	gogSvcLocF = flag.String(
		"google-svc-acc-key",
		"",
		"Google Clouds NLP JSON service account file, example: -google-svc-acc-key=\"~/google-svc-acc-key.json\"",
	)
)

func main() {
	log.Println("reading articles and entities ...")
	// Read
	articles, err := readArticles("./source/articles.csv")
	if err != nil {
		log.Fatal(err)
	}
	entities, err := readEntities("./source/entities.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Validate
	if len(*gogSvcLocF) == 0 {
		log.Fatal("missing google service account file")
	}
	if len(entities) == 0 {
		log.Fatal("missing entities")
	}
	log.Printf("len(entities)=%d", len(entities))

	texts := accumTexts(articles)
	if len(texts) == 0 {
		log.Fatal("missing texts")
	}
	log.Printf("len(texts)=%d", len(texts))

	log.Println("getting distances ...")
	nlpTok := nlp.NewNLPTokenizer(*gogSvcLocF, nlp.AutoLang)
	var wg sync.WaitGroup
	for _, entities := range entities {
		switch l := len(entities); {
		case l > 10:
			if err := scrape(texts, entities, nlpTok); err != nil {
				log.Fatal(err)
			}

			continue

		default:
			// Use goroutine strategy
		}

		wg.Add(1)

		go func(entities []string) {
			if err := scrape(texts, entities, nlpTok); err != nil {
				log.Fatal(err)
			}

			wg.Done()
		}(entities)
	}

	wg.Wait()
}

func scrape(texts, entities []string, tokenizer tokenize.Tokenizer) error {
	log.Printf("entities=%v", entities)

	// First entity is primary by definition
	entity := entities[0]
	log.Printf("entity=%s", entity)

	l := log.New(os.Stderr, "entity="+entity+",", 0)

	// Ignore articles without entity. This is a fuzzy search to spare the API
	temp := texts[:0]
	for _, text := range texts {
		if strings.Contains(text, entity) {
			temp = append(temp, text)
		}
	}
	texts = temp
	l.Printf("len(texts)=%d", len(texts))

	var (
		poS    = tokenize.ADJ | tokenize.ADP | tokenize.ADV | tokenize.CONJ | tokenize.DET | tokenize.NOUN | tokenize.NUM | tokenize.PRON | tokenize.PRT | tokenize.VERB
		source = assocentity.NewSource(entities, texts)
	)
	dists, err := assocentity.Distances(
		context.Background(),
		tokenizer,
		poS,
		source,
	)
	if err != nil {
		l.Fatal(err)
	}
	assocentity.Normalize(dists, assocentity.HumanReadableNormalizer)
	assocentity.Threshold(dists, 0.1)
	mean := assocentity.Mean(dists)

	l.Printf("len(mean)=%d", len(mean))

	if len(mean) == 0 {
		l.Print("no mean found, returning")
		return nil
	}

	// Convert to slice to make it sortable
	l.Println("converting to slice ...")
	type meanVal struct {
		dist float64
		tok  tokenize.Token
	}
	meanVals := make([]meanVal, 0)
	for tok, dist := range mean {
		// TODO: Whitelist ASCII
		meanVals = append(meanVals, meanVal{
			dist: dist,
			tok:  tok,
		})
	}

	// Sort by closest distance
	l.Println("sorting by pos and distance ...")
	sort.Slice(meanVals, func(i, j int) bool {
		if meanVals[i].tok.PoS != meanVals[j].tok.PoS {
			return meanVals[i].tok.PoS < meanVals[j].tok.PoS
		}
		return meanVals[i].dist < meanVals[j].dist
	})

	// Top 10 per pos
	l.Println("limit top 10 ...")
	type topMeanVal struct {
		Dist float64 `json:"distance"`
		PoS  string  `json:"pos"`
		Text string  `json:"text"`
	}
	topMeanVals := make([]topMeanVal, 0)
	poSCounter := make(map[tokenize.PoS]int)
	for _, meanVal := range meanVals {
		// Stop at 10 results per pos
		if poSCounter[meanVal.tok.PoS] >= 10 {
			continue
		}

		// Ignore unusual missmatches
		poS, ok := tokenize.PoSMapStr[meanVal.tok.PoS]
		if !ok {
			continue
		}
		topMeanVals = append(topMeanVals, topMeanVal{
			Dist: meanVal.dist,
			PoS:  poS,
			Text: meanVal.tok.Text,
		})

		poSCounter[meanVal.tok.PoS] += 1
	}
	l.Printf("len(topMeanVals)=%d", len(topMeanVals))

	// Write top 10 to disk
	l.Println("writing to disk ...")
	file, err := json.MarshalIndent(&topMeanVals, "", " ")
	if err != nil {
		l.Fatal(err)
	}
	name := url.QueryEscape(strings.ToLower(entity))
	path := filepath.Join("web", "public", name+".json")
	if err := os.WriteFile(path, file, 0600); err != nil {
		l.Fatal(err)
	}

	return nil
}

func readEntities(path string) (entities [][]string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entities = append(entities, strings.Split(scanner.Text(), ","))
	}
	return
}

func readArticles(path string) (articles [][]string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	articles, err = csvReader.ReadAll()
	return
}

func accumTexts(articles [][]string) (texts []string) {
	// For
	//	[
	//		[ID1, DATE1, LINK1, TITLE1, SUBTILE1, TEXT1],
	//		[ID2, DATE2, LINK2, TITLE2, SUBTILE2, TEXT2],
	//		...
	//	]
	for _, article := range articles[1:] { // Remove CSV header
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
				if len(text) > 0 {
					texts = append(texts, text)
				}
			}
		}
	}
	return
}
