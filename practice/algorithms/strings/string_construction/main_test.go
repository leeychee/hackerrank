package main

import "testing"

func TestConstruct(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"a", 1},
		{"abcd", 4},
		{"abab", 2},
	}
	for _, c := range cases {
		g := construct(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
