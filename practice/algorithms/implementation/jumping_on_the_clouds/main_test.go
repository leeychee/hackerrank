package main

import "testing"

var cases = []struct {
	s []int
	e int
}{
	{[]int{0, 0, 1, 0, 0, 1, 0}, 4},
	{[]int{0, 0, 0, 0, 1, 0}, 3},
}

func TestMinJumping(t *testing.T) {
	for _, c := range cases {
		g, err := minJumping(c.s)
		if err != nil {
			t.Errorf("%s", err)
		}
		if err == nil && c.e != g {
			t.Errorf("Expected: %d, Got: %d\n", c.e, g)
		}
	}
}
