package tickets

func (s *Service) ValidateOne(id string) (bool, string) {
	r := s.Validate(id)
	return r.Valid, r.Reason
}
