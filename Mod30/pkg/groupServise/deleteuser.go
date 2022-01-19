package groupServise

import (
	"errors"
)

func (s *Service) DeleteUser(ID int) (string, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Users[ID]; !ok {
		return "", errors.New("user not found")
	}

	name := s.Users[ID].Name
	delete(s.Users, ID)

	for key, _ := range s.Users {
		if _, ok := s.Users[key].Friends[ID]; ok {
			delete(s.Users[key].Friends, ID)
		}
	}

	return name, nil
}
