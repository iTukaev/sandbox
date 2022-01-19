package groupServise

func (s *Service) CreateUser(name string, age int) int {
	s.Lock()
	ID := s.NextUserID
	s.NextUserID++
	s.Unlock()

	s.Users[ID] = NewUser(ID, age, name)

	return ID
}