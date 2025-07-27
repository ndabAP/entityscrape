package cases

import (
	"path/filepath"

	"github.com/ndabAP/entityscrape/store"
)

var (
	GoogleCloudSvcAccountKey string
	SampleRate               uint64 = 100 // Default: 100 %
)

func NewStudy[samples, aggregated any](
	ident string,
	collector Collector[samples],
	aggregator Aggregator[samples, aggregated],
	reportor Reporter[aggregated],
) study[samples, aggregated] {
	file, err := store.NewFile(filepath.Join("reports", ident))
	if err != nil {
		panic(err.Error())
	}
	subjects := make(map[string]Analyses)
	return study[samples, aggregated]{
		Subjects:  subjects,
		collect:   collector,
		aggregate: aggregator,
		report:    reportor,
		store:     file,
	}
}
