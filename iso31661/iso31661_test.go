package iso31661

import (
	"testing"
)

func TestIs(t *testing.T) {
	cs := map[string]bool{
		"12": false,
		"AU": true,
		"SE": true,
	}
	for c, e := range cs {
		g := IsAlpha2(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	cs = map[string]bool{
		"123": false,
		"AUS": true,
		"SWE": true,
	}
	for c, e := range cs {
		g := IsAlpha3(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	cs = map[string]bool{
		"ABC": false,
		"036": true,
		"752": true,
	}
	for c, e := range cs {
		g := IsNumeric(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}

func TestName(t *testing.T) {
	cs := map[string]string{
		"12": "",
		"AU": "Australia",
		"SE": "Sweden",
	}
	for c, e := range cs {
		g := NameAlpha2(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	cs = map[string]string{
		"123": "",
		"AUS": "Australia",
		"SWE": "Sweden",
	}
	for c, e := range cs {
		g := NameAlpha3(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	cs = map[string]string{
		"ABC": "",
		"036": "Australia",
		"752": "Sweden",
	}
	for c, e := range cs {
		g := NameNumeric(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
