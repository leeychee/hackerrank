package main

import "testing"

func TestCompare(t *testing.T) {
	var cases = []struct {
		a, b   int
		ea, eb int
	}{
		{5, 3, 1, 0},
		{6, 6, 0, 0},
		{7, 10, 0, 1},
	}
	for _, c := range cases {
		ga, gb := compare(c.a, c.b)
		if ga != c.ea || gb != c.eb {
			t.Errorf("%d <> %d, Expected: %d, %d, Got: %d, %d", c.a, c.b, c.ea, c.eb, ga, gb)
		}
	}
}
