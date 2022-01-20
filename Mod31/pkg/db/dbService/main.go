package dbService

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty"`
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []string `json:"friends,omitempty"`
}

type Service struct {
	coll *mongo.Collection
}


func NewService() *Service {
	return &Service{
		coll: MongoDbCollection(),
	}
}