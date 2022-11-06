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

	// // CREATE
	// example_episode := episode.Episode{
	// 	Full_text: "I'm made of moon cheese",
	// 	Title:     "Episode Title",
	// 	Yt_id:     "owien23k"}

	// episode.Create(collection, example_episode)

	// // // READ
	// filter := bson.D{{"yt_id", "owien23k"}}
	// episode.Read(collection, filter)

	// // // UPDATE
	// episode.Update(collection, filter)

	// // DELETE
	// episode.Delete_all(collection)

	// SEARCH
	res := episode.Search(collection, "anything")
	for _, ep := range res {
		fmt.Println(ep.Title)
	}

	// Disconnect from MongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
