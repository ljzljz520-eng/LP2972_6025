package settlement

func (s *Service) ByCurrency(currency string) []Settlement {
	out := []Settlement{}
	for _, r := range s.records {
		if r.Currency == currency {
			out = append(out, r)
		}
	}
	return out
}
