package groupServise

import (
	"errors"
	"fmt"
)

func (s *service) DeleteUser(ID string) (string, error) {
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}

	name := s.Users[ID].Name
	delete(s.Users, ID)

	result := fmt.Sprintf("User ID: %d, name: %s was deleted", ID, name)
	return result, nil
}
