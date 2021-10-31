package service

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"sandbox/Mod31/pkg/entity"
)

func (cl *MongoClient) AgeUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		router := chi.URLParam(r, "userId")

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		u := entity.NewUpdateAge()
		if err := json.Unmarshal(content, u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		id, err := primitive.ObjectIDFromHex(router)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		collection := cl.Client.Client.Database("mod31").Collection("users")

		opts := options.FindOneAndUpdate().SetUpsert(false)
		filter := bson.D{{"_id", id}}
		update := bson.D{{"$set", bson.D{{"age", u.NewAge}}}}
		if err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Age of user: ID " + router + " updated successful"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}