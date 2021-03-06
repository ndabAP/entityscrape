package assoc

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ndabAP/assocentity/v8/tokenize"
	"github.com/ndabAP/entityscrape/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err       error
	client    *mongo.Client
	assoccoll *mongo.Collection
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

	assoccoll = client.Database("entityScrapeDB").Collection("associations")
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

	_, err = assoccoll.InsertMany(context.Background(), docs)
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

	if _, err := assoccoll.InsertOne(context.Background(), doc); err != nil {
		return err
	}

	return nil
}

// FindOne finds one
func FindOne(word, entity string) (*model.Assoc, error) {
	filter := bson.D{{Key: "word", Value: word}, {Key: "entitiy", Value: entity}}

	var w *model.Assoc
	err := assoccoll.FindOne(context.TODO(), filter).Decode(w)
	if err != nil {
		return w, err
	}

	return w, nil
}

// UpdateOne updates one
func UpdateOne(word, entity string, dist float64) error {
	filter := bson.D{{Key: "word", Value: word}, {Key: "entity", Value: entity}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "distance", Value: dist}}}}

	if _, err := assoccoll.UpdateOne(context.Background(), filter, update); err != nil {
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

var poSs = map[string]int{
	"adj":   tokenize.ADJ,
	"adp":   tokenize.ADP,
	"adv":   tokenize.ADV,
	"affix": tokenize.AFFIX,
	"any":   tokenize.ANY,
	"conj":  tokenize.CONJ,
	"det":   tokenize.DET,
	"noun":  tokenize.NOUN,
	"num":   tokenize.NUM,
	"pron":  tokenize.PRON,
	"prt":   tokenize.PRT,
	"punct": tokenize.PUNCT,
	"verb":  tokenize.VERB,
	"x":     tokenize.X,
}

// Aggregate aggregates
func Aggregate(entity, poS, from, to string) ([]Element, error) {
	var (
		err error
		f   time.Time
		t   time.Time
	)
	f, err = time.Parse("2006-01-02", from)
	if err != nil {
		return []Element{}, err
	}
	t, err = time.Parse("2006-01-02", to)
	if err != nil {
		return []Element{}, err
	}

	// Include starting today
	t.AddDate(0, 0, 1)

	// Include ending day
	t = t.AddDate(0, 0, 1)

	pipeline := []bson.M{
		bson.M{"$match": bson.M{
			"date": bson.M{
				"$gt": f.Format("2006-01-02T15:04:05Z"),
				"$lt": t.Format("2006-01-02T15:04:05Z"),
			},
			"entity": entity,
			"pos":    poSs[poS],
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
		bson.M{"$sort": bson.M{"count": 1}},
	}

	var cur *mongo.Cursor
	ctx := context.Background()

	cur, err = assoccoll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var aggregation []Element = []Element{}
	for cur.Next(ctx) {
		elem := &Element{}
		cur.Decode(elem)

		aggregation = append(aggregation, *elem)
	}

	cur.Close(ctx)

	return aggregation, nil
}

// Associations associations
func Associations(entity string) (int64, error) {
	filter := bson.D{{Key: "entity", Value: entity}}

	count, err := assoccoll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
