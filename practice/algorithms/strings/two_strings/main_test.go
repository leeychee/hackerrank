package main

import "testing"

func TestHaveCommonSubstring(t *testing.T) {
	var cases = []struct {
		s1, s2 string
		e      bool
	}{
		{"hello", "world", true},
		{"hi", "world", false},
	}
	for _, c := range cases {
		g := haveCommonSubstring(c.s1, c.s2)
		if g != c.e {
			t.Errorf("%q, %q, Expected: %t, Got: %t", c.s1, c.s2, c.e, g)
		}
	}
}
