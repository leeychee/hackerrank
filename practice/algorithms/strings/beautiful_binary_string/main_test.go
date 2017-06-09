package main

import "testing"

func TestMakeBeautiful(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"10", 0},
		{"110", 0},
		{"010", 1},
		{"01010", 1},
		{"0101010", 2},
		{"01100", 0},
		{"0100101010", 3},
	}
	for _, c := range cases {
		g := makeBeautiful(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
