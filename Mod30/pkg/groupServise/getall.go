package groupServise

import (
	"encoding/json"
)

func (s *Service) GetAll() ([]byte, error) {
	s.Lock()
	defer s.Unlock()

	body, err := json.Marshal(s.Users)
	if err != nil {
		return nil, err
	}
	return body, nil
}