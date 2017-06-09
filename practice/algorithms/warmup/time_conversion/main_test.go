package main

import "testing"

func TestTimeConv(t *testing.T) {
	var cases = []struct {
		s string
		e string
	}{
		{"07:05:45PM", "19:05:45"},
		{"12:40:22AM", "00:40:22"},
		{"12:45:54PM", "12:45:54"},
	}
	for _, c := range cases {
		g := timeConv(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %q, Got: %q", c.s, c.e, g)
		}
	}
}
