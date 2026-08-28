package api

import "testing"

func TestServiceHealth(t *testing.T) {
	if !New().Health() {
		t.Fatal()
	}
}
