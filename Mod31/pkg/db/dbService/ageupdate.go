package groupServise

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func (s *service) AgeUpdate(ID int, age int) (string, error) {
	//if _, ok := s.Users[ID]; !ok {
	//	return "", errors.New("user not found")
	//}
	//s.Users[ID].Age = age
	//
	//result := fmt.Sprintf("User's age ID: %d, name: %s was updated", ID, s.Users[ID].Name)
	//return result, nil
	collection := cl.Client.Client.Database("mod31").Collection("users")

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"age", u.NewAge}}}}
	if err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
