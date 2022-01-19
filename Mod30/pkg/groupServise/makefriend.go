package groupServise

import (
	"errors"
)

func (s *Service) MakeFriend(targetID int, sourceID int) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Users[targetID]; !ok {
		return errors.New("target user not found")
	}

	if _, ok := s.Users[sourceID]; !ok {
		return errors.New("source user not found")
	}

	if _, ok := s.Users[targetID].Friends[sourceID]; ok {
		return errors.New("users already friends")
	}

	s.Users[targetID].Friends[sourceID] = struct{}{}
	s.Users[sourceID].Friends[targetID] = struct{}{}
	return nil
}