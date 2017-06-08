package main

import "testing"

func TestCountWords(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"", 0},
		{"saveChangesInTheEditor", 5},
	}
	for _, c := range cases {
		g := countWords(c.s)
		if g != c.e {
			t.Errorf("Count(%q), Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
