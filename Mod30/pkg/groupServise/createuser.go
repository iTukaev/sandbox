package groupServise

import (
	"fmt"
	"sync"
)

func (s *service) CreateUser(name string, age int) (string, error) {
	mu := &sync.Mutex{}
	ID := lastUserId(s, mu) + 1

	s.Users[ID] = &User{
		Name: name,
		Age: age,
	}

	result := fmt.Sprintf("User ID: %d, name: %s was created", ID, name)
	return result, nil
}

func lastUserId(d *service, mu *sync.Mutex) int {
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