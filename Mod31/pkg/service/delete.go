package service

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"sandbox/Mod31/pkg/entity"
)

func (cl *MongoClient) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		user := entity.NewAddFriend()
		if err := json.Unmarshal(content, user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		collection := cl.Client.Client.Database("mod31").Collection("users")
		objectId, err := primitive.ObjectIDFromHex(user.TargetId)
		if err != nil{
			log.Fatalln(err)
		}

		err = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": objectId}).Err()
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user ID not found"))
			return
		}


		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User deleted"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

//convert id string to ObjectId
objectId, err := primitive.ObjectIDFromHex("5b9223c86486b341ea76910c")
if err != nil{
log.Println("Invalid id")
}

// find
result:= client.Database(database).Collection("user").FindOne(context.Background(), bson.M{"_id": objectId})
user := model.User{}
result.Decode(user)