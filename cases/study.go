package cases

import (
	"io/fs"
	"path/filepath"

	"github.com/ndabAP/entityscrape/store"
)

var (
	Corpus                   fs.FS
	GoogleCloudSvcAccountKey string
)

func NewStudy[samples, aggregated any](
	ident string,
	collect Collector[samples],
	aggregate Aggregator[samples, aggregated],
	report Reporter[aggregated],
) study[samples, aggregated] {
	store := store.NewFile(filepath.Join("cases", ident, "report"))
	subjects := make(map[string]Analyses)
	return study[samples, aggregated]{
		Subjects:  subjects,
		collect:   collect,
		aggregate: aggregate,
		report:    report,
		store:     store,
	}
}
