package main

import "testing"

func TestMakeAnagramByRemove(t *testing.T) {
	var cases = []struct {
		s1, s2 string
		e      int
	}{
		{"a", "a", 0},
		{"cde", "abc", 4},
		{"cbbe", "abbc", 2},
		{"cbbe", "abbbc", 3},
		{"cbbbe", "abbc", 3},
	}
	for _, c := range cases {
		g := makeAnagramByRemove(c.s1, c.s2)
		if g != c.e {
			t.Errorf("%q, %q, Expected: %d, Got: %d", c.s1, c.s2, c.e, g)
		}
	}
}
