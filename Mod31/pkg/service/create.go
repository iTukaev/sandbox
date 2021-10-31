package service

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"sandbox/Mod31/pkg/entity"
)

type MongoClient struct {
	Client *entity.MongoClient
}

func NewMongoClient(client *entity.MongoClient) *MongoClient {
	return &MongoClient{
		Client: client,
	}
}

func (cl *MongoClient) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		u := entity.NewUser()
		if err := json.Unmarshal(content, u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		collection := cl.Client.Client.Database("mod31").Collection("users")

		res, err := collection.InsertOne(context.TODO(), u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("user isn't created"))
			return
		}
		id := res.InsertedID.(primitive.ObjectID).Hex()

		w.Write([]byte("User ID: " + id + ", name: " + u.Name + " was created"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}