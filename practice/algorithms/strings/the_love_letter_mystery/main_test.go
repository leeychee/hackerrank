package main

import "testing"

func TestMakePalindrom(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"aba", 0},
		{"abb", 1},
		{"abc", 2},
		{"abcd", 4},
	}
	for _, c := range cases {
		g := makePalindrome(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
