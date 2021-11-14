package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) GetFriends(ID string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return "", err
	}

	var u User
	filter := bson.D{{"_id", objID}}
	opts := options.FindOne()

	if err = s.coll.FindOne(context.TODO(), filter, opts).Decode(&u); err != nil {
		return "", err
	}
	result := fmt.Sprintf("Friends of user ID: %s - %v", ID, u.Friends)
	return result, nil
}