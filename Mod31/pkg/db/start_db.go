package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func MongoDbCollection(ctx context.Context) (coll *mongo.Collection) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
		return
	}
	if err := client.Connect(context.TODO()); err != nil {
		panic(err)
		return
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
		return
	}

	coll = client.Database("mod31").Collection("users")
	log.Println("MongoDB started")
	<-ctx.Done()

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	log.Println("MongoDB stopped")
	return
}
