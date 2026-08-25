package storage

func (s *Store) Exists(id string) bool { _, e := s.Get(id); return e == nil }
