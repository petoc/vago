package common

import "testing"

func TestVal(t *testing.T) {
	cs := map[string][]bool{
		"12":   []bool{true, false},
		"123":  []bool{false, true},
		"1234": []bool{false, false},
		"A1":   []bool{true, false},
		"AB":   []bool{true, false},
		"ABC":  []bool{false, true},
		"ABCD": []bool{false, false},
	}
	for c, e := range cs {
		g := Val(c, 2, 2, nil)
		if g != e[0] {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
	for c, e := range cs {
		g := Val(c, 3, 3, nil)
		if g != e[1] {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
