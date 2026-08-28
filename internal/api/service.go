package api

import (
	"ticketverify/internal/audit"
	"ticketverify/internal/events"
	"ticketverify/internal/pricing"
	"ticketverify/internal/settlement"
	"ticketverify/internal/tickets"
)

type Service struct {
	Events     *events.Catalog
	Tickets    *tickets.Service
	Settlement *settlement.Service
	Audit      *audit.Log
}

func New() *Service {
	c := events.NewCatalog()
	return &Service{Events: c, Tickets: tickets.NewService(c), Settlement: settlement.New(pricing.Default()), Audit: audit.New()}
}
func (s *Service) Issue(id, event, holder, kind string) (tickets.Ticket, error) {
	t, e := s.Tickets.Issue(id, event, holder, kind)
	if e == nil {
		s.Audit.Record("issue-"+id, holder, "issue", id)
	}
	return t, e
}
func (s *Service) Validate(ids []string, size int) []tickets.Result {
	r := s.Tickets.ValidateBatch(ids, size)
	for _, x := range r {
		s.Audit.Record("val-"+x.TicketID, "gate", "validate", x.Reason)
	}
	return r
}
func (s *Service) Settle(id string, t tickets.Ticket) settlement.Settlement {
	s.Audit.Record("set-"+id, "finance", "settle", t.ID)
	return s.Settlement.Create(id, t)
}
