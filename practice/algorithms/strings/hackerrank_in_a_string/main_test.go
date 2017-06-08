package main

import "testing"

func TestIncludeHackerrank(t *testing.T) {
	var cases = []struct {
		s string
		e bool
	}{
		{"hereiamstackerrank", true},
		{"hackerworld", false},
	}
	for _, c := range cases {
		g := includeHackerrank(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %t, Got: %t", c.s, c.e, g)
		}
	}
}
