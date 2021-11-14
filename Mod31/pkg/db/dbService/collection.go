package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sandbox/Mod31/pkg/entity"
)

func MongoDbStart() (*entity.MongoClient, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	cl := entity.NewClient()
	cl.Client = client
	log.Println("MongoDB started")
	return cl, err
}
