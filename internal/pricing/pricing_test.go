package pricing

import "testing"

func TestPrice(t *testing.T) {
	p := Default()
	if p.Price("vip", 2) != 440 {
		t.Fatal()
	}
	if Bundle(p, "standard", 10) != 900 {
		t.Fatal()
	}
}
