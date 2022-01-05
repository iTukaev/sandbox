package groupServise

import (
	"errors"
	"fmt"
)

func (s *service) AgeUpdate(ID int, age int) (string, error) {
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}
	s.Users[ID].Age = age

	result := fmt.Sprintf("User's age ID: %d, name: %s was updated", ID, s.Users[ID].Name)
	return result, nil
}
