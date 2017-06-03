package main

import "testing"

func TestSpecialNum(t *testing.T) {
	var cases = []struct {
		k  int
		ts []int
		e  int
	}{
		{3, []int{4, 2, 6, 1, 10}, 4},
	}
	for _, c := range cases {
		g := specialNum(c.k, c.ts)
		if g != c.e {
			t.Errorf("%d, %v, Expected: %d, Got: %d\n", c.k, c.ts, c.e, g)
		}
	}
}
