package groupServise

import (
	"errors"
)

func (s *Service) AgeUpdate(ID int, age int) (string, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}

	s.Users[ID].Age = age
	return s.Users[ID].Name, nil
}
