package cli

import (
	"log"
	"math"

	"github.com/ndabAP/assocentity/v6/tokenize"
	"github.com/ndabAP/entityscrape/pkg/api"
	assocDB "github.com/ndabAP/entityscrape/pkg/db/assoc"
	newsDB "github.com/ndabAP/entityscrape/pkg/db/news"
	models "github.com/ndabAP/entityscrape/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// AssocEntitieser returns associated entities
type AssocEntitieser interface {
	AssocEntities(string, []string, *log.Logger) (map[tokenize.Token]float64, error)
}

var (
	entities = []string{
		"Angela Merkel",
		// "Barack Obama",
	}
	aliases = [][]string{
		{"Angela Dorothea Merkel", "Merkel"},
		// {"Barack Hussein Obama II", "Obama"},
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
			if ok, err := newsDB.Exists(n.ID); ok {
				if err != mongo.ErrNoDocuments && err != nil {
					return err
				}

				logger.Printf("news with id %s already exists, skipping", n.ID)

				continue
			} else {
				if err := newsDB.InsertOne(models.News{ID: n.ID}); err != nil {
					return err
				}
			}

			assocEntities, err := ae.AssocEntities(n.Text, append(aliases[idx], entity), logger)
			if err != nil {
				return err
			}

			logger.Printf("found %d associations", len(assocEntities))

			for word, dist := range assocEntities {
				if a, err := assocDB.FindOne(word.Token, entity); a == nil {
					if err != mongo.ErrNoDocuments && err != nil {
						return err
					}

					if err := assocDB.InsertOne(models.Assoc{
						Word:     word.Token,
						PoS:      word.PoS,
						Distance: dist,
						Entity:   entity,
					}); err != nil {
						return err
					}
				} else {
					dist := avg([]float64{a.Distance, dist})
					if err := assocDB.UpdateOne(word.Token, entity, dist); err != nil {
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
