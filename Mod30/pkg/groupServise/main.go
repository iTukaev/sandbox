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
	DeleteUser(ID int) (string, error)
	AgeUpdate(ID int, age int) (string, error)
	GetFriends(ID int) (string, error)
	MakeFriend(TargetID int, SourceID int) (string, error)
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