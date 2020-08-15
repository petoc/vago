package iso31662

import (
	"testing"
)

func TestIs(t *testing.T) {
	cs := map[string]bool{
		"12":     false,
		"AU-NSW": true,
		"AU-SA":  true,
		"SE-K":   true,
	}
	for c, e := range cs {
		g := Is(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}

func TestName(t *testing.T) {
	cs := map[string]string{
		"12":     "",
		"AU-NSW": "New South Wales, Australia",
		"AU-SA":  "South Australia, Australia",
		"SE-K":   "Blekinge län, Sweden",
	}
	for c, e := range cs {
		g := Name(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
