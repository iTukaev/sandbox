package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (s *service) DeleteUser(ID string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil{
		log.Println(err)
		return "", err
	}

	err = s.coll.FindOneAndDelete(context.TODO(), bson.M{"_id": objID}).Err()
	if err == mongo.ErrNoDocuments {
		log.Println(err)
		return "", err
	}

	filter := bson.D{{"friends", ID}}
	optsUpdate := options.FindOneAndUpdate().SetUpsert(false)
	update := bson.D{{"$pull", bson.D{{"friends", ID}}}}
	for {
		if err = s.coll.FindOneAndUpdate(context.TODO(), filter, update, optsUpdate).Err(); err != nil {
			break
		}
	}

	result := fmt.Sprintf("User ID: %s, was deleted", ID)
	return result, nil
}
