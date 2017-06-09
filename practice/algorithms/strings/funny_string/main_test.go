package main

import "testing"

func TestAbs(t *testing.T) {
	var cases = []struct {
		a, b byte
		e    int
	}{
		{'a', 'a', 0},
		{'a', 'b', 1},
		{'b', 'a', 1},
		{'a', 'z', 25},
	}
	for _, c := range cases {
		g := abs(c.a, c.b)
		if g != c.e {
			t.Errorf("abs(%s, %s), Expected: %d, Got: %d", string(c.a), string(c.b), c.e, g)
		}
	}
}

func TestFunny(t *testing.T) {
	var cases = []struct {
		s string
		e bool
	}{
		{"acxz", true},
		{"bcxz", false},
	}
	for _, c := range cases {
		g := funny(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %t, Got: %t", c.s, c.e, g)
		}
	}
}
