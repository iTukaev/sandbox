package groupServise

import (
	"errors"
	"strconv"
)

func (s *Service) GetFriends(ID int) ([]string, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Users[ID]; !ok {
		return nil, errors.New("user not found")
	}

	if len(s.Users[ID].Friends) == 0 {
		return nil, nil
	}

	friends := make([]string, 0, len(s.Users[ID].Friends))
	for key, _ := range s.Users[ID].Friends {
		friends = append(friends, strconv.Itoa(key))
	}

	return friends, nil
}
