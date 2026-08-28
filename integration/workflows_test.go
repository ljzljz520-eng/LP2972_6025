package integration

import (
	"testing"
	"ticketverify/internal/api"
	"time"
)

func TestWorkflowIssueAndValidate(t *testing.T) {
	s := api.New()
	s.Events.Create("e", "Final", "Arena", time.Now(), 3)
	if _, e := s.Issue("t", "e", "h", "vip"); e != nil {
		t.Fatal(e)
	}
	if !s.Validate([]string{"t"}, 1)[0].Valid {
		t.Fatal()
	}
}
func TestWorkflowSettlement(t *testing.T) {
	s := api.New()
	s.Events.Create("e", "F", "A", time.Now(), 3)
	tkt, _ := s.Issue("t", "e", "h", "standard")
	if s.Settle("s", tkt).Amount != 100 {
		t.Fatal()
	}
}
func TestWorkflowAudit(t *testing.T) {
	s := api.New()
	s.Audit.Record("1", "op", "scan", "t")
	if len(s.AuditEntries("op")) != 1 {
		t.Fatal()
	}
}
