package main

import "testing"

func TestMultiple(t *testing.T) {
	var cases = []struct {
		b byte
		n int
		e string
	}{
		{'#', 0, ""},
		{'#', 3, "###"},
	}
	for _, c := range cases {
		g := multiple(c.b, c.n)
		if g != c.e {
			t.Errorf("multiple(%s, %d), Expected: %q, Got: %q", string(c.b), c.n, c.e, g)
		}
	}
}

func TestStaircase(t *testing.T) {
	var cases = []struct {
		n int
		e string
	}{
		{1, `#`},
	}
	for _, c := range cases {
		g := staircase(c.n)
		if g != c.e {
			t.Errorf("staircase(%d), Expected:\n%s\nGot:\n%s", c.n, c.e, g)
		}
	}
}
