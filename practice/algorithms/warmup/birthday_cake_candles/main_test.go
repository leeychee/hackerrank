package main

import "testing"

func TestMaxNumCount(t *testing.T) {
	var cases = []struct {
		s []int
		e int
	}{
		{[]int{3, 2, 1, 3}, 2},
	}
	for _, c := range cases {
		g := maxNumCount(c.s)
		if g != c.e {
			t.Errorf("%v, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
