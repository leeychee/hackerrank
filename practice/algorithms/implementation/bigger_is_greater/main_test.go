package main

import "testing"

func TestBigger(t *testing.T) {
	var cases = []struct {
		s, e string
		err  error
	}{
		{"ab", "ba", nil},
		{"bb", "", ErrNoAnswer},
		{"hefg", "hegf", nil},
		{"dhck", "dhkc", nil},
		{"dkhc", "hcdk", nil},
	}
	for _, c := range cases {
		g, err := bigger(c.s)
		if err != c.err || g != c.e {
			t.Errorf("%q, Expected: %q, Got: %q\n", c.s, c.e, g)
		}
	}
}

func TestMinner(t *testing.T) {
	var cases = []struct {
		s, e string
	}{
		{"a", "a"},
		{"dkc", "cdk"},
	}
	for _, c := range cases {
		g := string(minner([]byte(c.s)))
		if g != c.e {
			t.Errorf("Expected: %s, Got: %s\n", c.e, g)
		}
	}
}
