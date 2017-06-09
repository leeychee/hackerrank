package main

import "testing"

var cases = []struct {
	s    []int
	mine int
	maxe int
}{
	{[]int{1, 2, 3, 4, 5}, 10, 14},
}

func TestMinSun(t *testing.T) {
	for _, c := range cases {
		g := minSum(c.s)
		if g != c.mine {
			t.Errorf("min(%v), Expected: %d, Got: %d", c.s, c.mine, g)
		}
	}
}

func TestMaxSun(t *testing.T) {
	for _, c := range cases {
		g := maxSum(c.s)
		if g != c.maxe {
			t.Errorf("max(%v), Expected: %d, Got: %d", c.s, c.maxe, g)
		}
	}
}
