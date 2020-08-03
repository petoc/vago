package iso31662

import (
	"testing"

	"github.com/petoc/isogo/iso"
)

func TestVal(t *testing.T) {
	cs := map[string]bool{
		"12":      false,
		"12-":     false,
		"12-1":    false,
		"12-12":   false,
		"12-123":  false,
		"A1":      false,
		"AB":      false,
		"AB-":     false,
		"AB-C":    true,
		"AB-CD":   true,
		"AB-CDEF": false,
		"AB-1":    true,
		"AB-12":   true,
		"AB-123":  true,
		"AB-1234": false,
	}
	for c, e := range cs {
		g := iso.Val(c, 4, 6, re)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}

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
		"AU-NSW": "New South Wales",
		"AU-SA":  "South Australia",
		"SE-K":   "Blekinge län",
	}
	for c, e := range cs {
		g := Name(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
