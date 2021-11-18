package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client
var err error
func MongoDbCollection() (coll *mongo.Collection) {
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
		return
	}
	if err = Client.Connect(context.TODO()); err != nil {
		panic(err)
		return
	}
	if err = Client.Ping(context.TODO(), nil); err != nil {
		panic(err)
		return
	}

	coll = Client.Database("mod31").Collection("users")
	log.Println("MongoDB started")

	return
}