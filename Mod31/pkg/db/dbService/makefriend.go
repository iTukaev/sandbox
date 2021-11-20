package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (s *service) MakeFriend(targetID string, sourceID string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(targetID)
	if err != nil {
		return "", err
	}

	//search target user in db
	filter := bson.D{{"_id", objID}}
	opts := options.FindOne()
	if err = s.coll.FindOne(context.TODO(), filter, opts).Err(); err != nil {
		log.Println(err)
		return "", err
	}

	//if err = findFriendInFriends(u.Friends, sourceID); err != nil {
	//	return "", err
	//}

	optsUpdate := options.FindOneAndUpdate().SetUpsert(false)
	update := bson.D{{"$addToSet", bson.D{{"friends", sourceID}}}}
	if err = s.coll.FindOneAndUpdate(context.TODO(), filter, update, optsUpdate).Err(); err != nil {
		log.Println(err)
		return "", err
	}
	result := fmt.Sprintf("User ID: %s now friend to user ID: %s", sourceID, targetID)
	return result, nil
}

//func findFriendInFriends(s []string, sourceID string) error {
//	for _, val := range s {
//		if val == sourceID {
//			return errors.New("these users are already friends")
//		}
//	}
//	return nil
//}