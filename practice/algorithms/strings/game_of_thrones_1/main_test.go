package main

import "testing"

func TestPalindromable(t *testing.T) {
	var cases = []struct {
		s string
		e bool
	}{
		{"aaabbbb", true},
		{"cdefghmnopqrstuvw", false},
		{"cdcdcdcdeeeef", true},
	}
	for _, c := range cases {
		g := palindromable(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %t, Got: %t", c.s, c.e, g)
		}
	}
}
