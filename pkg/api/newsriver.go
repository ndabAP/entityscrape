package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	apiKey string
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
	godotenv.Load()

	apiKey = os.Getenv("NEWSRIVER_API_KEY")
}

// Get gets
func Get(entity string) (News, error) {
	u, err := url.Parse(rootURL)
	if err != nil {
		return News{}, err
	}

	q := u.Query()

	query := fmt.Sprintf("%s:%s AND %s:EN", requestParams.title, strconv.Quote(entity), requestParams.language)
	q.Set(requestParams.query, query)
	q.Set(requestParams.sortBy, "discoverDate")
	q.Set(requestParams.sortOrder, "DESC")
	q.Set(requestParams.limit, "1")

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return News{}, err
	}

	req.Header.Set(requestParams.authorization, apiKey)

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
