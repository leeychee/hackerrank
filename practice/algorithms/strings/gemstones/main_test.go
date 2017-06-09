package main

import "testing"

func TestGem(t *testing.T) {
	var cases = []struct {
		ss []string
		e  int
	}{
		{[]string{}, 0},
		{[]string{"abcdde", "baccd", "eeabg"}, 2},
	}
	for _, c := range cases {
		g := gem(c.ss)
		if g != c.e {
			t.Errorf("%v, Expected: %d, Got: %d", c.ss, c.e, g)
		}
	}
}

func TestIntersection(t *testing.T) {
	var cases = []struct {
		s0, s1, e string
	}{
		{"abc", "cba", "cba"},
		{"abc", "cbac", "cba"},
		{"a", "b", ""},
	}
	for _, c := range cases {
		g := string(intersection([]byte(c.s0), []byte(c.s1)))
		if g != c.e {
			t.Errorf("%q inter %q, Expected: %q, Got: %q", c.s0, c.s1, c.e, g)
		}
	}
}
