package dbService

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Age int `json:"age" bson:"age"`
	Friends []string `json:"friends" bson:"friends"`
}

type Service struct {
	coll *mongo.Collection
}


func NewService() *Service {
	return &Service{
		coll: MongoDbCollection(),
	}
}