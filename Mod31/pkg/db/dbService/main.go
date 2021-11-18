package dbService

import (
	"bytes"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []string `json:"friends"`
}

type service struct {
	coll *mongo.Collection
}

type Service interface {
	CreateUser(name string, age int) (string, error)
	DeleteUser(ID string) (string, error)
	AgeUpdate(ID string, age int) (string, error)
	GetFriends(ID string) (string, error)
	MakeFriend(TargetID string, SourceID string) (string, error)
	GetAll() (*bytes.Buffer, error)
}

func NewService() Service {
	s := service{
		coll: MongoDbCollection(),
	}
	return &s
}

func (u *User) toString() string {
	return fmt.Sprintf("Name: %s, Age: %d, Friends: %v", u.Name, u.Age, u.Friends)
}