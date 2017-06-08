package main

import "testing"

func TestValid(t *testing.T) {
	var cases = []struct {
		bs   string
		a, b byte
		e    bool
	}{
		{"beabeefeab", 'a', 'b', true},
		{"beabeefeab", 'e', 'b', false},
	}
	for _, c := range cases {
		g := valid([]byte(c.bs), c.a, c.b)
		if g != c.e {
			t.Errorf("%s %s %s, Expected: %t, Got: %t", c.bs, string(c.a), string(c.b), c.e, g)
		}
	}
}

func TestMaxCharacters(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"beabeefeab", 5},
	}
	for _, c := range cases {
		g := maxTwoCharacters([]byte(c.s))
		if g != c.e {
			t.Errorf("%s, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
