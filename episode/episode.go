package episode

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Episode struct {
	Full_text string `bson:"full_text"`
	Title     string `bson:"title"`
	Yt_id     string `bson:"yt_id"`
}

// CREATE
func Create(collection *mongo.Collection, episode Episode) {

	insertResult, err := collection.InsertOne(context.TODO(), episode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// READ
func Read(collection *mongo.Collection, filter bson.D) Episode {
	var result Episode

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result
}

// UPDATE
func Update(collection *mongo.Collection, filter bson.D) {

	update := bson.D{
		{"$set", bson.D{
			{"full_text", "Cry all your tears, your update is here"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// DELETE
func Delete_all(collection *mongo.Collection) {
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the collection\n", deleteResult.DeletedCount)
}

// SEARCH
func Search(collection *mongo.Collection, searchterm string) []Episode {
	searchStage := bson.D{{"$search", bson.D{{"text", bson.D{{"path", "full_text"}, {"query", searchterm}}}}}}
	projectStage := bson.D{{"$project", bson.D{{"Yt_id", 1}, {"full_text", 1}, {"title", 1}, {"_id", 0}}}}
	// specify the amount of time the operation can run on the server
	opts := options.Aggregate().SetMaxTime(5 * time.Second)
	// run pipeline
	cursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{searchStage, projectStage}, opts)
	if err != nil {
		panic(err)
	}
	// print results
	var search_results []Episode
	for cursor.Next(context.TODO()) {
		var result Episode
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		search_results = append(search_results, result)
	}
	return search_results
}
