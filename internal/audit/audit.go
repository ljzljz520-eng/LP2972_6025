package audit

import (
	"strings"
	"time"
)

type Entry struct {
	ID, Actor, Action, Target string
	At                        time.Time
}

const AuditEntry = "AuditEntry"

type Log struct{ entries []Entry }

func New() *Log { return &Log{entries: []Entry{}} }
func (l *Log) Record(id, actor, action, target string) Entry {
	e := Entry{id, actor, action, target, time.Now().UTC()}
	l.entries = append(l.entries, e)
	return e
}
func (l *Log) Query(actor string) []Entry {
	out := []Entry{}
	for _, e := range l.entries {
		if actor == "" || e.Actor == actor {
			out = append(out, e)
		}
	}
	return out
}
func (l *Log) Search(term string) []Entry {
	out := []Entry{}
	for _, e := range l.entries {
		if strings.Contains(e.Action, term) || strings.Contains(e.Target, term) {
			out = append(out, e)
		}
	}
	return out
}
