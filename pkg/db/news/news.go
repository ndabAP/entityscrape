package news

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ndabAP/entityscrape/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err        error
	client     *mongo.Client
	collection *mongo.Collection
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	mongoConnURI := os.Getenv("MONGODB_CONNECTION_STRING")

	client, err = mongo.NewClient(options.Client().ApplyURI(mongoConnURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("entityScrapeDB").Collection("news")
}

// InsertOne inserts one
func InsertOne(news model.News) error {
	doc := bson.M{
		"_id":          news.ID,
		"entity":       news.Entity,
		"associations": news.Associations,
	}

	if _, err := collection.InsertOne(context.Background(), doc); err != nil {
		return err
	}

	return nil
}

// ReplaceOne replaces one
func ReplaceOne(news model.News) error {
	filter := bson.D{{Key: "_id", Value: news.ID}}
	doc := bson.M{
		"_id":          news.ID,
		"entity":       news.Entity,
		"associations": news.Associations,
	}

	if _, err := collection.ReplaceOne(context.Background(), filter, doc); err != nil {
		return err
	}

	return nil
}

// Exists checks if exists
func Exists(id, entity string, associations bool) (bool, error) {
	filter := bson.D{{Key: "_id", Value: id}, {Key: "entity", Value: entity}, {Key: "associations", Value: associations}}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}

// Count counts
func Count(entity string) (int64, error) {
	filter := bson.D{{Key: "entity", Value: entity}}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
