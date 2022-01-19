package groupServise

import (
	"fmt"
	"sync"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Friends map[int]struct{} `json:"friends"`
}

func NewUser(ID, age int, name string) *User {
	return &User{
		ID: ID,
		Name: name,
		Age: age,
		Friends: make(map[int]struct{}),
	}
}

type Service struct {
	sync.Mutex
	NextUserID int
	Users map[int]*User
}

func NewService() *Service {
	return &Service{
		NextUserID: 0,
		Users: make(map[int]*User),
	}
}

func (u *User) toString() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Friends: %v", u.ID, u.Name, u.Age, u.Friends)
}