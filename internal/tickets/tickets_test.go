package tickets

import (
	"testing"
	"ticketverify/internal/events"
	"time"
)

func TestIssueValidate(t *testing.T) {
	c := events.NewCatalog()
	c.Create("e", "x", "v", time.Now(), 2)
	s := NewService(c)
	if _, e := s.Issue("t", "e", "h", "standard"); e != nil {
		t.Fatal(e)
	}
	if !s.Validate("t").Valid {
		t.Fatal("invalid")
	}
}
