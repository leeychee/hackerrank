package main

import "testing"
import "reflect"

func TestKaprekar(t *testing.T) {
	var cases = []struct {
		p, q int
		eKar []int
		eErr error
	}{
		{1, 100, []int{1, 9, 45, 55, 99}, nil},
	}
	for _, c := range cases {
		g, ge := kaprekar(c.p, c.q)
		if ge != c.eErr || !reflect.DeepEqual(g, c.eKar) {
			t.Errorf("Expected: %v, Got: %v\n", c.eKar, g)
		}
	}
}

func TestIsKaprekar(t *testing.T) {
	var cases = []struct {
		x int
		e bool
	}{
		{1, true},
		{10, false},
		{45, true},
	}
	for _, c := range cases {
		g := isKarpreka(c.x)
		if g != c.e {
			t.Errorf("%d, Expected: %t, Got: %t\n", c.x, c.e, g)
		}
	}
}

func TestLenInt(t *testing.T) {
	var cases = []struct {
		x, e int
	}{
		{9, 1},
		{10, 2},
		{987, 3},
		{48596, 5},
	}
	for _, c := range cases {
		g := lenInt(c.x)
		if g != c.e {
			t.Errorf("Expected: %d, Got: %d\n", c.e, g)
		}
	}
}
