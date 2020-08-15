package iso4217

import (
	"testing"
)

func TestIs(t *testing.T) {
	cs := map[string]bool{
		"12":  false,
		"123": false,
		"AU":  false,
		"AUD": true,
		"SE":  false,
		"SEK": true,
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
		"12":  "",
		"123": "",
		"AU":  "",
		"AUD": "Australian Dollar",
		"SE":  "",
		"SEK": "Swedish Krona",
	}
	for c, e := range cs {
		g := Name(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
