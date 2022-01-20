package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func (s *Service) AgeUpdate(ID string, age int) error {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", bson.D{{"age", age}}}}
	if err := s.coll.FindOneAndUpdate(ctx, filter, update).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
