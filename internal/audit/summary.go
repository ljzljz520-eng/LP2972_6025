package audit

func (l *Log) Count() int { return len(l.entries) }
func (l *Log) Actions() map[string]int {
	m := map[string]int{}
	for _, e := range l.entries {
		m[e.Action]++
	}
	return m
}
