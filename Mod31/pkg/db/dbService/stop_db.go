package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func MongoDbStop(client *mongo.Client)  {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	log.Println("MongoDB stopped")
}
