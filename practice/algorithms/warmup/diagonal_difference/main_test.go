package main

import "testing"

func TestDistance(t *testing.T) {
	var cases = []struct {
		a, b int
		e    int
	}{
		{1, 3, 2},
		{3, 1, 2},
		{1, 1, 0},
	}
	for _, c := range cases {
		g := distance(c.a, c.b)
		if g != c.e {
			t.Errorf("distance(%d, %d), Expected: %d, Got: %d", c.a, c.b, c.e, g)
		}
	}
}
