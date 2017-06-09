package main

import "testing"

func TestPlusMinus(t *testing.T) {
	var cases = []struct {
		s       []int
		p, z, n int
	}{
		{[]int{-4, 3, -9, 0, 4, 1}, 3, 1, 2},
	}
	for _, c := range cases {
		gp, gz, gn := plusMinus(c.s)
		if gp != c.p || gz != c.z || gn != c.n {
			t.Errorf("%v, Expected: g%d, z%d, n%d, Got: g%d, z%d, n%d", c.s, c.p, c.z, c.z, gp, gz, gn)
		}
	}
}
