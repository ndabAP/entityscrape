package srv

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ndabAP/entityscrape/pkg/db/assoc"
	"github.com/ndabAP/entityscrape/pkg/db/news"
)

// Entities entities
func Entities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}

	entity := r.URL.Query().Get("entity")
	if entity == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	aggregation, err := assoc.Aggregate(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	marshaled, err := json.Marshal(aggregation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaled)
}

// News news
func News(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}

	entity := r.URL.Query().Get("entity")
	if entity == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	count, err := news.Count(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.FormatInt(count, 10)))
}
