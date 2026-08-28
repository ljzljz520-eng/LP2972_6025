package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
	"ticketverify/internal/tickets"
)

var bucket = []byte("tickets")

const ValidationRecord = "ValidationRecord"

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	e = db.Update(func(tx *bbolt.Tx) error { _, x := tx.CreateBucketIfNotExists(bucket); return x })
	if e != nil {
		db.Close()
		return nil, e
	}
	return &Store{db}, nil
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Put(t tickets.Ticket) error {
	b, e := json.Marshal(t)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(t.ID), b) })
}
func (s *Store) Get(id string) (tickets.Ticket, error) {
	var t tickets.Ticket
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucket).Get([]byte(id))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, &t)
	})
	return t, e
}
func (s *Store) List() ([]tickets.Ticket, error) {
	out := []tickets.Ticket{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(_, v []byte) error {
			var t tickets.Ticket
			if e := json.Unmarshal(v, &t); e != nil {
				return e
			}
			out = append(out, t)
			return nil
		})
	})
	return out, e
}
func (s *Store) Reopen(path string) (*Store, error) { s.Close(); return Open(path) }
