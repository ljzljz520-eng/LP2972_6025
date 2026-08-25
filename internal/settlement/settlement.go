package settlement

import (
	"ticketverify/internal/pricing"
	"ticketverify/internal/tickets"
	"time"
)

type Settlement struct {
	ID, TicketID, Currency string
	Amount                 int
	CreatedAt              time.Time
}
type Service struct {
	policy  pricing.Policy
	records map[string]Settlement
}

func New(p pricing.Policy) *Service { return &Service{policy: p, records: map[string]Settlement{}} }
func (s *Service) Create(id string, t tickets.Ticket) Settlement {
	amount := s.policy.Price(t.Kind, 1)
	r := Settlement{ID: id, TicketID: t.ID, Currency: "CNY", Amount: amount, CreatedAt: time.Now().UTC()}
	s.records[id] = r
	return r
}
func (s *Service) Get(id string) (Settlement, bool) { r, ok := s.records[id]; return r, ok }
func (s *Service) Total() int {
	n := 0
	for _, r := range s.records {
		n += r.Amount
	}
	return n
}
