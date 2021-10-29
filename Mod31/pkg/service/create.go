package service

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"sandbox/Mod31/pkg/db"
	"sandbox/Mod31/pkg/entity"
	"strconv"
)

type CreateUser struct {
	department *entity.Department
}

func NewCreateUser(department *entity.Department) *CreateUser {
	return &CreateUser{
		department: department,
	}
}

func (d *CreateUser) Create(w http.ResponseWriter, r *http.Request) {
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

		collection := db.Client.Database("mod31").Collection("users")

		u.Id = lastUserId(collection) + 1

		_, err = collection.InsertOne(context.TODO(), u)
		if err != nil {
			log.Fatalln(err)
		}

		w.Write([]byte("User ID: " + strconv.Itoa(u.Id) + " name: " + u.Name + " was created"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func lastUserId(c *mongo.Collection) int {
	var u *entity.User
	opts := options.FindOne().SetMax(bson.D{{"id", 1}})
	res := c.FindOne(context.TODO(), bson.D{{"id",1}}, opts)
	if err := res.Decode(u); err != nil {
		log.Fatalln(err)
	}

	return u.Id
}