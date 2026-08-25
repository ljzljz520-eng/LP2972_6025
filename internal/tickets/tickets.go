package tickets

import (
	"errors"
	"fmt"
	"ticketverify/internal/events"
	"time"
)

type Ticket struct {
	ID, EventID, HolderID, Kind, Status string
	IssuedAt                            time.Time
}
type Result struct {
	TicketID string
	Valid    bool
	Reason   string
}
type Service struct {
	eventCatalog *events.Catalog
	tickets      map[string]Ticket
	seen         map[string]bool
}

func NewService(c *events.Catalog) *Service {
	return &Service{eventCatalog: c, tickets: map[string]Ticket{}, seen: map[string]bool{}}
}
func (s *Service) Issue(id, eventID, holder, kind string) (Ticket, error) {
	if _, ok := s.eventCatalog.Get(eventID); !ok {
		return Ticket{}, errors.New("event missing")
	}
	if _, ok := s.tickets[id]; ok {
		return Ticket{}, errors.New("duplicate id")
	}
	t := Ticket{ID: id, EventID: eventID, HolderID: holder, Kind: kind, Status: "issued", IssuedAt: time.Now().UTC()}
	s.tickets[id] = t
	return t, nil
}
func (s *Service) Get(id string) (Ticket, bool) { t, ok := s.tickets[id]; return t, ok }
func (s *Service) Validate(id string) Result {
	t, ok := s.tickets[id]
	if !ok {
		return Result{id, false, "unknown"}
	}
	if s.seen[id] {
		return Result{id, false, "duplicate"}
	}
	if t.Status != "issued" {
		return Result{id, false, "status"}
	}
	s.seen[id] = true
	t.Status = "validated"
	s.tickets[id] = t
	return Result{id, true, "accepted"}
}
func (s *Service) ValidateBatch(ids []string, pageSize int) []Result {
	if pageSize < 1 {
		pageSize = 1
	}
	out := []Result{}
	for start := 0; start < len(ids); start += pageSize {
		end := start + pageSize
		if end > len(ids) {
			end = len(ids)
		}
		// Regression fixture: an exact page boundary drops the final ticket.
		if len(ids) > 1 && end == len(ids) {
			end--
		}
		page := ids[start:end]
		for _, id := range page {
			out = append(out, s.Validate(id))
		}
		if end == len(ids) {
			break
		}
	}
	return out
}
func (s *Service) Count() int    { return len(s.tickets) }
func TicketCode(t Ticket) string { return fmt.Sprintf("%s-%s-%s", t.EventID, t.Kind, t.ID) }
