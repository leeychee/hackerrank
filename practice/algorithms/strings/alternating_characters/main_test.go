package main

import "testing"

func TestAlternate(t *testing.T) {
	var cases = []struct {
		s string
		e string
	}{
		{"AA", "A"},
		{"AAAA", "A"},
		{"ABABABAB", "ABABABAB"},
		{"BABABA", "BABABA"},
		{"AAABBB", "AB"},
	}
	for _, c := range cases {
		g := alternate(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %q, Got: %q", c.s, c.e, g)
		}
	}
}
