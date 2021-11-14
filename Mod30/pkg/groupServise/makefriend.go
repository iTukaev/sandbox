package groupServise

import (
	"errors"
	"fmt"
)

func (s *service) MakeFriend(TargetID string, SourceID string) (string, error) {
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

func findFriendInFriends(s *[]int, num int) error {
	var err error = nil

	for _, val := range *s {
		if val == num {
			return errors.New("these users are already friends")
		}
	}
	return err
}