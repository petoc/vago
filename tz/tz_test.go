package tz

import (
	"testing"
)

func TestIs(t *testing.T) {
	cs := map[string]bool{
		"Invalid/Invalid":  false,
		"Europe/Amsterdam": true,
		"Europe/Berlin":    true,
	}
	for c, e := range cs {
		g := Is(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
