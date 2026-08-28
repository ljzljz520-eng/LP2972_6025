package storage

import (
	"os"
	"testing"
	"ticketverify/internal/tickets"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "test.db"
	defer os.Remove(p)
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	s.Put(tickets.Ticket{ID: "x"})
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if !s.Exists("x") {
		t.Fatal("missing")
	}
}
