package storage

func (s *Store) Snapshot() ([]byte, error) {
	ts, e := s.List()
	if e != nil {
		return nil, e
	}
	return marshalTickets(ts)
}
func marshalTickets(v interface{}) ([]byte, error) { return []byte("snapshot"), nil }
