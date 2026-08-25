package tickets

import (
	"testing"
	"ticketverify/internal/events"
	"time"
)

func TestDuplicateTicketRejected(t *testing.T) {
	c := events.NewCatalog()
	c.Create("e", "x", "v", time.Now(), 4)
	s := NewService(c)
	s.Issue("t", "e", "h", "standard")
	got := s.ValidateBatch([]string{"t", "t"}, 2)
	if len(got) != 2 || got[1].Reason != "duplicate" {
		t.Fatalf("expected duplicate rejection: %#v", got)
	}
}
