package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/ndabAP/assocentity/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	prose "gopkg.in/jdkato/prose.v2"
)

const (
	unicodeSmallA   = 97
	unicodeSmallZ   = 122
	unicodeCapitalA = 65
	unicodeCapitalZ = 90
)

var adjectives []string

func init() {
	file, err := os.Open("adjectives.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		adjectives = append(adjectives, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 10 * time.Second,
	})

	// On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	// Print link
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// 	// Visit link found on page
	// 	// Only those links are visited which are in AllowedDomains
	// 	c.Visit(e.Request.AbsoluteURL(link))
	// })

	c.OnHTML(".article-wrapper", func(e *colly.HTMLElement) {
		p := e.ChildText("p")

		weighting, err := associate(p, []string{"Trump", "Donald Trump", "D. Trump", "D. J. Trump", "Donald John Trump"})
		if err != nil {
			log.Fatal(err)
		}
		weightingAdjectives := keepAdjectives(weighting)
		err = insert(weightingAdjectives, "Donald John Trump", "adjective")
		if err != nil {
			log.Fatal(err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://eu.usatoday.com/story/opinion/2019/04/16/trump-show-features-acting-officials-key-posts-editorial-debate/3486558002/")
}

func associate(text string, entity []string) (weighting map[string]float64, err error) {
	weighting, err = assocentity.Make(text, entity, func(text string) ([]string, error) {
		tokenizedText, err := tokenizer(text)
		if err != nil {
			return nil, err
		}

		return tokenizedText, nil
	})

	if err != nil {
		return nil, err
	}

	return weighting, nil
}

func tokenizer(text string) ([]string, error) {
	document, err := prose.NewDocument(string(text))
	if err != nil {
		return nil, err
	}

	var tokenizedText []string
	for _, token := range document.Tokens() {
		ok := true
		for _, r := range token.Text {
			// Only allow latin alphabet
			switch {
			case r < unicodeCapitalA:
				ok = false
				break
			case r > unicodeSmallZ:
				ok = false
				break
			case r > unicodeCapitalZ && r < unicodeSmallA:
				ok = false
			}
		}

		if ok {
			tokenizedText = append(tokenizedText, token.Text)
		}
	}

	return tokenizedText, nil
}

func keepAdjectives(weighting map[string]float64) map[string]float64 {
	for word := range weighting {
		if !isInSlice(word, adjectives) {
			delete(weighting, word)
		}
	}

	return weighting
}

func insert(weighting map[string]float64, entity, wordType string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	collection := client.Database("entityScrape").Collection("weighting")
	var documents []interface{}
	for word, weight := range weighting {
		documents = append(documents, bson.M{"entity": entity, "word": word, "weight": weight, "type": wordType})
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		return err
	}

	return nil
}

func isInSlice(el string, slice []string) bool {
	for _, e := range slice {
		if e == el {
			return true
		}
	}

	return false
}
