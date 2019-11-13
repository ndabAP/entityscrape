package assoc

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ndabAP/assocentity/v6/tokenize"
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
func InsertMany(assocs []model.Assoc) error {
	var docs []interface{}
	for _, assoc := range assocs {
		doc := bson.M{
			"distance": assoc.Distance,
			"pos":      assoc.PoS,
			"entity":   assoc.Entity,
			"word":     assoc.Word,
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
func InsertOne(assoc model.Assoc) error {
	doc := bson.M{
		"date":     assoc.Date,
		"distance": assoc.Distance,
		"entity":   assoc.Entity,
		"pos":      assoc.PoS,
		"word":     assoc.Word,
	}

	if _, err := collection.InsertOne(context.Background(), doc); err != nil {
		return err
	}

	return nil
}

// FindOne finds one
func FindOne(word, entity string) (*model.Assoc, error) {
	filter := bson.D{{Key: "word", Value: word}, {Key: "entitiy", Value: entity}}

	var w *model.Assoc
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

// Element is element
type Element struct {
	Count    float64 `json:"count"`
	Distance float64 `json:"distance"`
	Word     string  `json:"word"`
}

// Aggregate aggregates
func Aggregate(entity string) ([]Element, error) {
	pipeline := []bson.M{
		bson.M{"$match": bson.M{
			"entity": entity,
			"pos":    tokenize.ADJ,
		}},
		bson.M{"$group": bson.M{
			"_id":      "$word",
			"count":    bson.M{"$sum": 1},
			"distance": bson.M{"$push": "$distance"},
		}},
		bson.M{"$sort": bson.M{"count": -1}},
		bson.M{"$limit": 10},
		bson.M{"$project": bson.M{
			"_id":      0,
			"word":     "$_id",
			"count":    "$count",
			"distance": bson.M{"$avg": "$distance"},
		}},
		bson.M{"$sort": bson.M{"distance": 1}},
	}

	var cur *mongo.Cursor
	ctx := context.Background()

	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var aggregation []Element
	for cur.Next(ctx) {
		elem := &Element{}
		cur.Decode(elem)

		aggregation = append(aggregation, *elem)
	}

	cur.Close(ctx)

	return aggregation, nil
}
