package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
var Client *mongo.Client

func MongoDbStart() {
	Client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalln(err)
	}
	if err := Client.Connect(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	if err := Client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
	}

	log.Println("MongoDB started")
}
