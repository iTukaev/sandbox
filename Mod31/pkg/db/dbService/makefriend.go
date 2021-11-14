package dbService

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) MakeFriend(targetID string, sourceID string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(targetID)
	if err != nil {
		return "", err
	}

	var u User
	filter := bson.D{{"_id", objID}}
	opts := options.FindOne()
	if err = s.coll.FindOne(context.TODO(), filter, opts).Decode(&u); err != nil {
		return "", err
	}
	if err = findFriendInFriends(u.Friends, sourceID); err != nil {
		return "", err
	}
	optsUpdate := options.FindOneAndUpdate().SetUpsert(false)
	update := bson.D{{"$set", bson.D{{"friends", sourceID}}}}
	s.coll.FindOneAndUpdate(context.TODO(), filter, update, optsUpdate)
	result := fmt.Sprintf("User ID: %s and user ID: %s already friends", targetID, sourceID)
	return result, nil
}

func findFriendInFriends(s []string, sourceID string) error {
	for _, val := range s {
		if val == sourceID {
			return errors.New("these users are already friends")
		}
	}
	return nil
}