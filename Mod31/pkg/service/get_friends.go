package service

import (
	"context"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
)

func (cl *MongoClient) Friends(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		router := chi.URLParam(r, "userId")

		id, err := primitive.ObjectIDFromHex(router)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		collection := cl.Client.Client.Database("mod31").Collection("users")

		filter := bson.D{{"_id", id}}
		opts := options.FindOne()
		collection.FindOne(context.TODO(), filter, opts)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Friends of user " + d.department.Users[uNum].Name + ": " + allFriendsOfUser))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}