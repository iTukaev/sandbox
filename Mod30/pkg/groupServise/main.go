package groupServise

import (
	"errors"
	"fmt"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Friends []int `json:"friends"`
}

type Service struct {
	Users map[int]*User
}

func NewService() *Service {
	s := &Service{
		make(map[int]*User),
	}
	return s
}

func (s *Service) CreateUser(name string, age int) (string, error) {
	mu := &sync.Mutex{}
	ID := lastUserId(s, mu) + 1

	s.Users[ID] = &User{
		Name: name,
		Age: age,
	}

	result := fmt.Sprintf("User ID: %d, name: %s was created", ID, name)
	return result, nil
}

func (s *Service) DeleteUser(ID int) (string, error) {
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}

	name := s.Users[ID].Name
	delete(s.Users, ID)

	result := fmt.Sprintf("User ID: %d, name: %s was deleted", ID, name)
	return result, nil
}

func (s *Service) UpdateUserAge(ID int, age int) (string, error) {
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}
	s.Users[ID].Age = age

	result := fmt.Sprintf("User's age ID: %d, name: %s was updated", ID, s.Users[ID].Name)
	return result, nil
}

func (s *Service) GetFriends(ID int) (string, error) {
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}

	result := ""

	if len(s.Users[ID].Friends) == 0 {
		result = fmt.Sprintf("User %s has not friends", s.Users[ID].Name)
		return result, nil
	}

	result = fmt.Sprintf("Friends of user ID: %d -\n", ID)
	for _, val := range s.Users[ID].Friends {
		result += fmt.Sprintf("\tID: %d\n", val)
	}

	return result, nil
}

func (s *Service) MakeFriend(TargetID int, SourceID int) (string, error) {
	if _, ok := s.Users[TargetID]; !ok {
		return "", errors.New("target user not found")
	}

	if _, ok := s.Users[SourceID]; !ok {
		return "", errors.New("source user not found")
	}

	if err := findFriendInFriends(&s.Users[TargetID].Friends, SourceID); err != nil {
		return "", err
	}

	s.Users[TargetID].Friends = append(s.Users[TargetID].Friends, SourceID)
	result := fmt.Sprintf("User ID: %d added to friend's list to user ID: %d", SourceID, TargetID)
	return result, nil
}

func lastUserId(d *Service, mu *sync.Mutex) int {
	mu.Lock()
	lastId := 0
	for key,  _ := range d.Users {
		if key > lastId {
			lastId = key
		}
	}
	mu.Unlock()
	return lastId
}

func findFriendInFriends(s *[]int, num int) error {
	var err error = nil

	for _, val := range *s {
		if val == num {
			return errors.New("these users are already friends")
		}
	}
	return err
}