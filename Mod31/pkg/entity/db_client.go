package entity

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient struct {
	Client *mongo.Client
}

func NewClient() *MongoClient {
	return new(MongoClient)
}