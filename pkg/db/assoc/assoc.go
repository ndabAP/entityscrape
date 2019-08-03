package assoc

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

	collection = client.Database("entityScrapeDB").Collection("associations")
}

// InsertMany inserts many documents at once
func InsertMany(accocs []model.Accoc) error {
	var docs []interface{}
	for _, accoc := range accocs {
		doc := bson.M{
			"distance": accoc.Distance,
			"entity":   accoc.Entity,
			"word":     accoc.Word,
		}

		docs = append(docs, doc)
	}

	_, err = collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}

	return nil
}

// InsertOne inserts one
func InsertOne(accoc model.Accoc) error {
	doc := bson.M{
		"distance": accoc.Distance,
		"entity":   accoc.Entity,
		"word":     accoc.Word,
	}

	if _, err := collection.InsertOne(context.Background(), doc); err != nil {
		return err
	}

	return nil
}

// FindOne finds one
func FindOne(word, entity string) (*model.Accoc, error) {
	filter := bson.D{{Key: "word", Value: word}, {Key: "entitiy", Value: entity}}

	var w *model.Accoc
	err := collection.FindOne(context.TODO(), filter).Decode(w)
	if err != nil {
		return w, err
	}

	return w, nil
}

// UpdateOne updates one
func UpdateOne(word, entity string, dist float64) error {
	filter := bson.D{{Key: "word", Value: word}, {Key: "entity", Value: entity}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "distance", Value: dist}}}}

	if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
		return err
	}

	return nil
}