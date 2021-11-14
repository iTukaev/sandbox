package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) DeleteUser(ID string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil{
		return "", err
	}

	err = s.coll.FindOneAndDelete(context.TODO(), bson.M{"_id": objectId}).Err()
	if err == mongo.ErrNoDocuments {
		return "", err
	}

	result := fmt.Sprintf("User ID: %s, was deleted", ID)
	return result, nil
}
