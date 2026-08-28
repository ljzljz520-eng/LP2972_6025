package events

import (
	"testing"
	"time"
)

func TestEventLifecycle(t *testing.T) {
	c := NewCatalog()
	c.Create("e", "x", "v", time.Now(), 1)
	if !c.Available("e", 0) {
		t.Fatal()
	}
	c.Deactivate("e")
	if c.Available("e", 0) {
		t.Fatal()
	}
}
