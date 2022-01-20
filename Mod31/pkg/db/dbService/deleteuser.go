package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func (s *Service) DeleteUser(ID string) error {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil{
		log.Println(err)
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	err = s.coll.FindOneAndDelete(ctx, bson.M{"_id": objID}).Err()
	if err == mongo.ErrNoDocuments {
		log.Println(err)
		return err
	}

	filter := bson.D{{"friends", ID}}
	update := bson.D{{"$pull", bson.D{{"friends", ID}}}}
	if _, err = s.coll.UpdateMany(ctx, filter, update); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
