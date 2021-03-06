package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	apiKey   string
	apiLimit string
)

const (
	rootURL = "https://api.newsriver.io/v2/search"
)

var requestParams = struct {
	language, title, query, sortBy, sortOrder, limit, authorization string
}{
	"language", "title", "query", "sortBy", "sortOrder", "limit", "Authorization",
}

// News represents news
type News []struct {
	ID   string
	Text string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	apiKey = os.Getenv("NEWSRIVER_API_KEY")
	apiLimit = os.Getenv("NEWSRIVER_API_LIMIT")
}

// Get gets
func Get(entity string, logger *log.Logger) (News, error) {
	u, err := url.Parse(rootURL)
	if err != nil {
		return News{}, err
	}

	q := u.Query()

	query := fmt.Sprintf("%s:%s AND %s:EN", requestParams.title, strconv.Quote(entity), requestParams.language)
	q.Set(requestParams.query, query)
	q.Set(requestParams.sortBy, "discoverDate")
	q.Set(requestParams.sortOrder, "DESC")
	q.Set(requestParams.limit, apiLimit)

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return News{}, err
	}

	req.Header.Set(requestParams.authorization, apiKey)

	log.Printf("requesting news api with url: %s", u.String())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return News{}, err
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return News{}, err
	}

	var news News
	json.Unmarshal(b, &news)

	return news, nil
}
