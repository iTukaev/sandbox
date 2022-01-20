package dbService

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func (s *Service) MakeFriend(TargetID string, SourceID string) error {
	targetObjID, err := primitive.ObjectIDFromHex(TargetID)
	if err != nil {
		return err
	}
	sourceObjID, err := primitive.ObjectIDFromHex(SourceID)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	targetFilter := bson.D{{"_id", targetObjID}}
	if err = s.coll.FindOne(ctx, targetFilter).Err(); err != nil {
		log.Println(err)
		return fmt.Errorf("user %s not found %w", targetObjID, err)
	}

	sourceFilter := bson.D{{"_id", sourceObjID}}
	if err = s.coll.FindOne(ctx, sourceFilter).Err(); err != nil {
		log.Println(err)
		return fmt.Errorf("user %s not found %w", sourceObjID, err)
	}

	targetUpdate := bson.D{{"$addToSet", bson.D{{"friends", SourceID}}}}
	if err = s.coll.FindOneAndUpdate(ctx, targetFilter, targetUpdate).Err(); err != nil {
		log.Println(err)
		return err
	}

	sourceUpdate := bson.D{{"$addToSet", bson.D{{"friends", TargetID}}}}
	if err = s.coll.FindOneAndUpdate(ctx, sourceFilter, sourceUpdate).Err(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}