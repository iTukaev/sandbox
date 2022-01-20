package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func (s *Service) CreateUser(name string, age int) (string, error) {
	u := User{
		Name: name,
		Age: age,
		Friends: []string{},
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	res, err := s.coll.InsertOne(ctx, u)
	if err != nil {
		log.Printf("User %s creating error %v",name, err)
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
