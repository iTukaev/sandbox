package groupServise

import (
	"bytes"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []int `json:"friends"`
}

type service struct {
	Users map[int]*User
}

type Service interface {
	CreateUser(name string, age int) (string, error)
	DeleteUser(ID string) (string, error)
	AgeUpdate(ID string, age int) (string, error)
	GetFriends(ID string) (string, error)
	MakeFriend(TargetID string, SourceID string) (string, error)
	GetAll() *bytes.Buffer
}

func NewService() Service {
	s := service{
		make(map[int]*User),
	}
	return &s
}

func (u *User) toString() string {
	return fmt.Sprintf("Name: %s, Age: %d, Friends: %v", u.Name, u.Age, u.Friends)
}