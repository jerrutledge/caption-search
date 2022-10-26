package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jerrutledge/caption-search/episode"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// please note that an environment variable must be set for this code to successfully connect to the db

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB.")

	// Panic if disconnected
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("caption-search").Collection("episodes")

	example_episode := episode.Episode{
		Full_text: "This is the full text",
		Title:     "Episode Title",
		Yt_id:     "owien23k"}

	insertResult, err := collection.InsertOne(context.TODO(), example_episode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"

	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	// Disconnect from MongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
