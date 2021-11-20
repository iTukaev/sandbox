package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func (s *service) CreateUser(name string, age int) (string, error) {
	u := User{
		Name: name,
		Age: age,
		Friends: []string{},
	}

	res, err := s.coll.InsertOne(context.TODO(), u)
	if err != nil {
		log.Println(err)
		return "", err
	}
	ID := res.InsertedID.(primitive.ObjectID).Hex()

	result := fmt.Sprintf("User ID: %s, name: %s was created", ID, name)
	return result, nil
}
