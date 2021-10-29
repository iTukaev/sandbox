package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Employ struct {
	Id int `json:"_id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []int `json:"friends"`
}

func (e *Employ) New() *Employ {
	return new(Employ)
}

func main()  {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalln(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
	}

	log.Println("MongoDB started")

	collection := client.Database("mod31").Collection("users")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()
	//one := Employ{Id: 1,Name: "Tara",Age: 24, Friends: []int{1, 4, 5}}
	//insertResult, err := collection.InsertOne(context.TODO(), one)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Println("Inserted a single document: ", insertResult.InsertedID)

	//filter := bson.D{{"name", "Tara"}}

	//update := bson.D{
	//	{"$set", bson.D{
	//		{"age", 45},
	//	}},
	//}

	//updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	f := collection.Indexes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", &f)

}
