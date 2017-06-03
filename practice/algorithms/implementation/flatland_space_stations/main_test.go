package main

import "testing"

func TestMaxDistance(t *testing.T) {
	var cases = []struct {
		n  int
		cs []int
		e  int
	}{
		{5, []int{0, 4}, 2},
		{6, []int{0, 1, 2, 3, 4, 5}, 0},
		{20, []int{13, 1, 11, 10, 6}, 6},
	}
	for _, c := range cases {
		g := maxDistance(c.n, c.cs)
		if g != c.e {
			t.Errorf("%d, %v, Expected: %d, Got: %d\n", c.n, c.cs, c.e, g)
		}
	}
}
