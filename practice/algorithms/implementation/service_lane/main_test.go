package main

import "testing"

func TestPcType(t *testing.T) {
	var cases = []struct {
		l lane
		s *service
		e int
	}{
		{lane{2, 3, 1, 2, 3, 2, 3, 3}, &service{0, 3}, 1},
		{lane{2, 3, 1, 2, 3, 2, 3, 3}, &service{4, 6}, 2},
		{lane{2, 3, 1, 2, 3, 2, 3, 3}, &service{6, 7}, 3},
		{lane{2, 3, 1, 2, 3, 2, 3, 3}, &service{3, 5}, 2},
		{lane{2, 3, 1, 2, 3, 2, 3, 3}, &service{0, 7}, 1},
	}
	for _, c := range cases {
		g := pctype(c.l, c.s)
		if g != c.e {
			t.Errorf("%v, %v, Expected: %d, Got: %d\n", c.l, c.s, c.e, g)
		}
	}
}
