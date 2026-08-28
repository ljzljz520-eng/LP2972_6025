package settlement

import (
	"testing"
	"ticketverify/internal/pricing"
	"ticketverify/internal/tickets"
)

func TestSettlementTotal(t *testing.T) {
	s := New(pricing.Default())
	s.Create("1", tickets.Ticket{ID: "t", Kind: "standard"})
	if s.Total() != 100 {
		t.Fatal()
	}
}
