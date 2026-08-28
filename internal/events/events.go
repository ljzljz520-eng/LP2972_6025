package events

import "time"

type Event struct {
	ID, Name, Venue string
	StartsAt        time.Time
	Capacity        int
	Active          bool
}
type Catalog struct{ items map[string]Event }

func NewCatalog() *Catalog { return &Catalog{items: map[string]Event{}} }
func (c *Catalog) Create(id, name, venue string, start time.Time, capacity int) Event {
	e := Event{ID: id, Name: name, Venue: venue, StartsAt: start, Capacity: capacity, Active: true}
	c.items[id] = e
	return e
}
func (c *Catalog) Get(id string) (Event, bool) { e, ok := c.items[id]; return e, ok }
func (c *Catalog) Deactivate(id string) bool {
	e, ok := c.items[id]
	if !ok {
		return false
	}
	e.Active = false
	c.items[id] = e
	return true
}
func (c *Catalog) List() []Event {
	out := make([]Event, 0, len(c.items))
	for _, e := range c.items {
		out = append(out, e)
	}
	return out
}
