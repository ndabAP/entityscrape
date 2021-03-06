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
	poS := r.URL.Query().Get("part-of-speech")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	if entity == "" || poS == "" || from == "" || to == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	aggregation, err := assoc.Aggregate(entity, poS, from, to)
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

// Associations associations
func Associations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}

	entity := r.URL.Query().Get("entity")
	if entity == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	count, err := assoc.Associations(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.FormatInt(count, 10)))
}

// List list
func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	list := []string{
		"Angela Merkel",
		"El Chapo",
		"Elon Musk",
		"Donald Trump",
		"Greta Thunberg",
		"Kim Kardashian",
		"Vladimir Putin",
		"Volodymyr Zelensky",
		"Xi Jinping",
	}

	marshaled, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(marshaled)
}
