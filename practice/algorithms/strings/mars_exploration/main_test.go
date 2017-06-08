package main

import "testing"

func TestChanged(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"SOSSPSSQSSOR", 3},
		{"SOSSOT", 1},
	}
	for _, c := range cases {
		g := changed(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
