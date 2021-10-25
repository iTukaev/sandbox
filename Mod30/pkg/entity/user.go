package entity

import "encoding/json"

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []int `json:"friends"`
}

func NewUser() *User {
	return new(User)
}

func (u *User) ToString() string {
	b, _ := json.Marshal(u)
	return string(b)
}
