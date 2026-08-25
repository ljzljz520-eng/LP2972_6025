package events

func (c *Catalog) IsSoldOut(id string, sold int) bool { return c.Remaining(id, sold) == 0 }
