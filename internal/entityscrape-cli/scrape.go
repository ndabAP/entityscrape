package cli

import (
	"log"

	"github.com/gocolly/colly"
)

// Make makes
func Make(entity, url string, aliases []string) error {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnHTML(".article-wrapper", func(e *colly.HTMLElement) {
		p := e.ChildText("p")

		aliases = append([]string{entity}, aliases...)
		weighting, err := weighting(p, aliases)
		if err != nil {
			log.Fatal(err)
		}

		weightingAdjectives := keepAdjectives(weighting)
		log.Println(len(weightingAdjectives), "adjectives found")

		err = insert(weightingAdjectives, entity, "adjective")
		if err != nil {
			return err
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)
}
