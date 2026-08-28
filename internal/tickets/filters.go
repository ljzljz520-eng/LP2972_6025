package tickets

func (s *Service) ByEvent(eventID string) []Ticket {
	out := []Ticket{}
	for _, t := range s.tickets {
		if t.EventID == eventID {
			out = append(out, t)
		}
	}
	return out
}
func (s *Service) ByHolder(holder string) []Ticket {
	out := []Ticket{}
	for _, t := range s.tickets {
		if t.HolderID == holder {
			out = append(out, t)
		}
	}
	return out
}
