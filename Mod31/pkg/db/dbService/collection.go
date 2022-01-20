package dbService

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client
var err error

type Config struct {
	MongoServer string `yaml:"mongo_server"`
	Database string `yaml:"database"`
	Collection string `yaml:"collection"`
}

func MongoDbCollection() (coll *mongo.Collection) {
	cfg := new(Config)
	if err = cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		log.Fatalf("config reading error: %v", err)
	}

	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + cfg.MongoServer))
	if err != nil {
		log.Fatalf("mongo new client error: %v", err)
	}

	if err = Client.Connect(context.TODO()); err != nil {
		log.Fatalf("mongo connect error: %v", err)
	}

	if err = Client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("mongo ping error: %v", err)
	}

	coll = Client.Database(cfg.Database).Collection(cfg.Collection)
	log.Println("MongoDB started")

	return
}