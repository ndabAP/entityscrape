package cli

import (
	"log"
	"math"

	"github.com/ndabAP/entityscrape/pkg/api"
	db "github.com/ndabAP/entityscrape/pkg/db/assoc"
	models "github.com/ndabAP/entityscrape/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// AssocEntitieser returns associated entities
type AssocEntitieser interface {
	AssocEntities(string, []string, *log.Logger) (map[string]float64, error)
}

var (
	entities = []string{"Angela Merkel", "Barack Obama"}
	aliases  = [][]string{
		{"Angela Dorothea Merkel", "Merkel", "A. Merkel", "Angela M."},
		{"Barack Hussein Obama II", "B. Obama", "Obama"},
	}
)

// Do does
func Do(ae AssocEntitieser, logger *log.Logger) error {
	for idx, entity := range entities {
		logger.Printf("getting news for entity: %s", entity)

		news, err := api.Get(entity, logger)
		if err != nil {
			return err
		}

		logger.Printf("found %d news", len(news))

		for _, n := range news {
			assocEntities, err := ae.AssocEntities(n.Text, aliases[idx], logger)
			if err != nil {
				return err
			}

			logger.Printf("found %d associations", len(assocEntities))

			for word, dist := range assocEntities {
				if a, err := db.FindOne(word, entity); a == nil {
					if err != mongo.ErrNoDocuments && err != nil {
						return err
					}

					if err := db.InsertOne(models.Accoc{
						Word:     word,
						Distance: dist,
						Entity:   entity,
					}); err != nil {
						return err
					}
				} else {
					dist := avg([]float64{a.Distance, dist})
					if err := db.UpdateOne(word, entity, dist); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// Returns the average of a float slice
func avg(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return round(total / float64(len(xs)))
}

// Rounds to nearest 0.5
func round(x float64) float64 {
	return math.Round(x/0.5) * 0.5
}
