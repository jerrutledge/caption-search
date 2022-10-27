package episode

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Episode struct {
	Full_text string
	Title     string
	Yt_id     string
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
