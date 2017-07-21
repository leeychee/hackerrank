package main

import "testing"
import "reflect"

func TestExpand(t *testing.T) {
	var cases = []struct {
		x, n int
		e    [][]int
	}{
		{10, 2, [][]int{{1, 3}}},
		{100, 2, [][]int{{1, 3, 4, 5, 7}, {6, 8}, {10}}},
	}
	for _, c := range cases {
		g := expand(c.x, c.n)
		if !reflect.DeepEqual(g, c.e) {
			t.Errorf("x: %d, n: %d, E: %v, G: %v", c.x, c.n, c.e, g)
		}
	}
}

func TestPower(t *testing.T) {
	var cases = []struct {
		x, y int
		e    int
	}{
		{0, 2, 0},
		{1, 2, 1},
		{2, 2, 4},
		{10, 2, 100},
	}
	for _, c := range cases {
		g := power(c.x, c.y)
		if g != c.e {
			t.Errorf("%d^%d, Expected: %d, Got: %d", c.x, c.y, c.e, g)
		}
	}
}
