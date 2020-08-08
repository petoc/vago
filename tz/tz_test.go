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

func TestNames(t *testing.T) {
	cs := map[string][]string{
		"12": nil,
		"DE": []string{"Europe/Berlin", "Europe/Busingen"},
		"NL": []string{"Europe/Amsterdam"},
	}
	for c, e := range cs {
		g := Names(c)
		if len(g) != len(e) {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
		for i, v := range g {
			if v != e[i] {
				t.Errorf("%s: expected %v, got %v", c, e, g)
				return
			}
		}
	}
}
