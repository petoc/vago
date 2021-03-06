package iso31663

import (
	"testing"
)

func TestIs(t *testing.T) {
	cs := map[string]bool{
		"ABCD": false,
		"BQAQ": true,
		"TPTL": true,
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
		"ABCD": "",
		"BQAQ": "British Antarctic Territory",
		"TPTL": "East Timor",
	}
	for c, e := range cs {
		g := Name(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
