package main

import "testing"

func TestFeast(t *testing.T) {
	var cases = []struct {
		t *trip
		e int
	}{
		{&trip{10, 2, 5}, 6},
		{&trip{12, 4, 4}, 3},
		{&trip{6, 2, 2}, 5},
		{&trip{857, 2, 43}, 438},
	}
	for _, c := range cases {
		g := c.t.feast()
		if c.e != g {
			t.Errorf("%s, Expected: %d, Got: %d\n", c.t, c.e, g)
		}
	}
}
