package main

import "testing"

var cases = []struct {
	a byte
	s string
	n int
	e int
}{
	{'a', "a", 100000000, 100000000},
	{'a', "aba", 10, 7},
}

func TestCountRepeatedString(t *testing.T) {
	for _, c := range cases {
		g := countRepeatedString(c.a, c.s, c.n)
		if g != c.e {
			t.Errorf("Expected: %d, Got: %d\n", c.e, g)
		}
	}
}
