package main

import "testing"

func TestIsAnagram(t *testing.T) {
	var cases = []struct {
		s1, s2 string
		e      bool
	}{
		{"abc", "cba", true},
		{"abc", "bba", false},
	}
	for _, c := range cases {
		g := isAnagram(c.s1, c.s2)
		if g != c.e {
			t.Errorf("%q, %q, Expected: %t, Got: %t", c.s1, c.s2, c.e, g)
		}
	}
}

func TestMakeAnagram(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"aaabbb", 3},
		{"ab", 1},
		{"abc", -1},
		{"mnop", 2},
		{"xyyx", 0},
		{"xaxbbbxx", 1},
		{"xzxbbbxx", 1},
	}
	for _, c := range cases {
		g := makeAnagram(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
