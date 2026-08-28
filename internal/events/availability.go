package events

func (c *Catalog) Available(id string, sold int) bool {
	e, ok := c.Get(id)
	if !ok {
		return false
	}
	if !e.Active {
		return false
	}
	return sold < e.Capacity
}
func (c *Catalog) Remaining(id string, sold int) int {
	e, ok := c.Get(id)
	if !ok {
		return 0
	}
	r := e.Capacity - sold
	if r < 0 {
		return 0
	}
	return r
}
