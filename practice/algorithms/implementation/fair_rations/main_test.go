package main

import "testing"

func TestMakeEven(t *testing.T) {
	var cases = []struct {
		bs   []int
		eErr error
		e    int
	}{
		{[]int{1, 2}, ErrUnableMakeEven, 0},
		{[]int{2, 3, 4, 5, 6}, nil, 4},
	}
	for _, c := range cases {
		g, err := makeEven(c.bs)
		if err != c.eErr || g != c.e {
			t.Errorf("%v, Expected: %d, Got: %d\n", c.bs, c.e, g)
		}
	}
}
