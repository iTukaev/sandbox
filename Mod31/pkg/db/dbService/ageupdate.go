package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) AgeUpdate(ID string, age int) (string, error) {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return "", err
	}

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", bson.D{{"age", age}}}}
	if err := s.coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("User's ID: %s age updated", ID), nil
}
