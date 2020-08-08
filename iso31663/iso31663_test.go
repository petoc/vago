package iso31663

import (
	"testing"

	"github.com/petoc/vago/common"
)

func TestVal(t *testing.T) {
	cs := map[string]bool{
		"12":    false,
		"123":   false,
		"1234":  false,
		"A1":    false,
		"AB":    false,
		"ABC":   false,
		"ABC1":  false,
		"ABCD":  true,
		"ABCDE": false,
	}
	for c, e := range cs {
		g := common.Val(c, 4, 4, re)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}

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
