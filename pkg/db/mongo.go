package repository

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err          error
	client       *mongo.Client
	collection   *mongo.Collection
	mongoConnURI = os.Getenv("MONGODB_CONNECTION_STRING")
)

func init() {
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoConnURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("entityScrapeDB").Collection("weighting")
}

// InsertMany inserts many documents at once
func InsertMany(docs []interface{}) error {
	_, err = collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}

	return nil
}
