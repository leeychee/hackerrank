package main

import "testing"

func TestSimpleSum(t *testing.T) {
	var cases = []struct {
		s []int
		e int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3}, 6},
	}
	for _, c := range cases {
		g := simpleSum(c.s)
		if g != c.e {
			t.Errorf("%v, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
