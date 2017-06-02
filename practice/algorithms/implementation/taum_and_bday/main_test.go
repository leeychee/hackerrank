package main

import "testing"

var cases = []struct {
	p *plan
	e int
}{
	{&plan{10, 10, 1, 1, 1}, 20},
	{&plan{5, 9, 2, 3, 4}, 37},
	{&plan{3, 6, 9, 1, 1}, 12},
	{&plan{7, 7, 4, 2, 1}, 35},
	{&plan{3, 3, 1, 9, 2}, 12},
}

func TestPlanSpend(t *testing.T) {
	for _, c := range cases {
		g := c.p.spend()
		if g != c.e {
			t.Errorf("%v, Expected: %d, Got: %d\n", c.p, c.e, g)
		}
	}
}
