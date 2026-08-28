package api

func (s *Service) Health() bool {
	return s != nil && s.Events != nil && s.Tickets != nil && s.Audit != nil
}
