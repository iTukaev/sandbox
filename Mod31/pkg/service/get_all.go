package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func (cl *MongoClient) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)

		collection := cl.Client.Client.Database("mod31").Collection("users")

		documents, err := collection.Find(context.TODO(), bson.M{}, options.Find())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		for documents.Next(context.TODO()) {
			////this variant get all users without canonicals
			//byteString, err := bson.MarshalExtJSON(documents.Current, false, false)
			//
			//if err != nil {
			//	log.Fatalln(err)
			//	continue
			//}
			//byteString = append(byteString, '\n')
			//w.Write(byteString)
			currentUser := documents.Current.String() + "\n"
			w.Write([]byte(currentUser))
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}