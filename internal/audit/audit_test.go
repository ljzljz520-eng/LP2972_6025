package audit

import "testing"

func TestAuditQuery(t *testing.T) {
	l := New()
	l.Record("1", "a", "validate", "t")
	if len(l.Query("a")) != 1 {
		t.Fatal()
	}
}
