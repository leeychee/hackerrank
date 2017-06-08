package main

import "testing"

func TestSuperReduce(t *testing.T) {
	var cases = []struct {
		s, e string
	}{
		{"aa", ""},
		{"baab", ""},
		{"aaabccddd", "abd"},
	}
	for _, c := range cases {
		g := superReduce(c.s)
		if g != c.e {
			t.Errorf("%s, Expected: %s, Got: %s", c.s, c.e, g)
		}
	}
}
