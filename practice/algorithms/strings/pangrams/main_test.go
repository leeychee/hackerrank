package main

import "testing"

func TestIspangram(t *testing.T) {
	var cases = []struct {
		s string
		e bool
	}{
		{"The quick brown fox jumps over the lazy dog", true},
		{"We promptly judged antique ivory buckles for the next prize", true},
		{"We promptly judged antique ivory buckles for the prize", false},
	}
	for _, c := range cases {
		g := isPangram(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %t, Got: %t", c.s, c.e, g)
		}
	}
}
