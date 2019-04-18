package cli

import (
	"log"

	"github.com/gocolly/colly"
)

// Make makes
func Make(entity, url string, aliases []string) error {
	errorc := make(chan error, 1)

	go func() {
		c := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		)

		c.OnHTML("body", func(e *colly.HTMLElement) {
			log.Println("body", e.Text)
		})

		c.OnHTML(".article-wrapper", func(e *colly.HTMLElement) {
			log.Println(".article-wrapper", e.Text)

			p := e.ChildText("p")

			aliases = append([]string{entity}, aliases...)
			weighting, err := weighting(p, aliases)
			if err != nil {
				errorc <- err
			}

			weightingAdjectives := keepAdjectives(weighting)
			log.Println(len(weightingAdjectives), "adjectives found")

			err = insert(weightingAdjectives, entity, "adjective")
			if err != nil {
				errorc <- err
			}
		})

		c.OnRequest(func(r *colly.Request) {
			log.Println("visiting", r.URL.String())
		})

		c.OnResponse(func(r *colly.Response) {
			log.Println("status", r.StatusCode)
		})

		c.OnError(func(r *colly.Response, err error) {
			errorc <- err
		})

		c.OnScraped(func(r *colly.Response) {
			errorc <- nil
		})

		c.Visit(url)
	}()

	err := <-errorc
	if err != nil {
		return err
	}

	return nil
}
