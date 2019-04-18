package cli

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insert(weighting map[string]float64, entity, wordType string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	collection := client.Database("entityScrape").Collection("weighting")
	var documents []interface{}
	for word, weight := range weighting {
		documents = append(documents, bson.M{"entity": entity, "word": word, "weight": weight, "type": wordType})
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		return err
	}

	return nil
}
