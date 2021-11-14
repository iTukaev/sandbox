package groupServise

import (
	"errors"
	"fmt"
)

func (s *service) GetFriends(ID string) (string, error) {
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
