package iso31661

import (
	"testing"

	"github.com/petoc/vago/common"
)

func TestVal(t *testing.T) {
	cs := map[string][]bool{
		"A1":   []bool{false, false, false},
		"AB":   []bool{true, false, false},
		"ABC":  []bool{false, true, false},
		"ABCD": []bool{false, false, false},
		"12":   []bool{false, false, false},
		"123":  []bool{false, false, true},
		"1234": []bool{false, false, false},
	}
	for c, e := range cs {
		g := common.Val(c, 2, 2, reA)
		if g != e[0] {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	for c, e := range cs {
		g := common.Val(c, 3, 3, reA)
		if g != e[1] {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	for c, e := range cs {
		g := common.Val(c, 3, 3, reN)
		if g != e[2] {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}

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
