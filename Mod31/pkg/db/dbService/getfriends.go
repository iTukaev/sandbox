package dbService

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *Service) GetFriends(ID string) ([]string, error) {
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	u := new(User)
	filter := bson.D{{"_id", objID}}

	if err = s.coll.FindOne(ctx, filter).Decode(u); err != nil {
		return nil, err
	}

	return u.Friends, nil
}