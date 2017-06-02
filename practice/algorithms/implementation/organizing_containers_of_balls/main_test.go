package main

import "testing"

var cases = []struct {
	m matrix
	e bool
}{
	{matrix{{1, 1}, {1, 1}}, true},
	{matrix{{0, 2}, {1, 1}}, false},
}

func TestMatrixCanSwapToPure(t *testing.T) {
	for _, c := range cases {
		g := c.m.canSwapToPure()
		if g != c.e {
			t.Errorf("\n%s, Expected: %t, Got: %t\n", c.m, c.e, g)
		}
	}
}
