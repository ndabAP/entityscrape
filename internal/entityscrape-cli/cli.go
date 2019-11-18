package cli

import (
	"log"
	"strings"
	"time"

	"github.com/ndabAP/assocentity/v7/tokenize"
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
		"Elon Musk",
		"Donald Trump",
		"Vladimir Putin",
		"Xi Jinping",
	}
	aliases = [][]string{
		{"Angela Dorothea Merkel", "Merkel"},
		{"Elon Reeve Musk", "Musk"},
		{"Trump"},
		{"Vladimir Vladimirovich Putin", "Putin"},
		{"Jinping"},
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
			if ok, err := newsDB.Exists(n.ID, entity, true); ok {
				if err != mongo.ErrNoDocuments && err != nil {
					return err
				}

				logger.Printf("news with id %s and entity %s already exists, skipping", n.ID, entity)

				continue
			} else {
				if ok, err := newsDB.Exists(n.ID, entity, false); !ok {
					if err != mongo.ErrNoDocuments && err != nil {
						return err
					}

					if err := newsDB.InsertOne(models.News{ID: n.ID, Entity: entity, Associations: false}); err != nil {
						return err
					}
				}
			}

			if strings.TrimSpace(n.Text) == "" {
				logger.Printf("empty text, skipping")

				continue
			}

			assocEntities, err := ae.AssocEntities(n.Text, append(aliases[idx], entity), logger)
			if err != nil {
				return err
			}

			newsDB.ReplaceOne(models.News{
				Associations: len(assocEntities) != 0,
				Entity:       entity,
				ID:           n.ID,
			})

			logger.Printf("found %d associations", len(assocEntities))

			for word, dist := range assocEntities {
				if err := assocDB.InsertOne(models.Assoc{
					Date:     time.Now().UTC().Format(models.DateFormat),
					Distance: dist,
					Entity:   entity,
					PoS:      word.PoS,
					Word:     word.Token,
				}); err != nil {
					return err
				}
			}

			time.Sleep(time.Second * 10)
		}

		time.Sleep(time.Second * 5)
	}

	return nil
}
