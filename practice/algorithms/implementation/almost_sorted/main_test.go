package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	var cases = []struct {
		s []int
		e []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
	}
	for _, c := range cases {
		s0 := make([]int, len(c.s))
		copy(s0, c.s)
		reverse(s0)
		if !reflect.DeepEqual(s0, c.e) {
			t.Errorf("reverse(%v), Expected: %v, Got: %v", c.s, c.e, s0)
		}
	}
}
func TestSortMethod(t *testing.T) {
	var cases = []struct {
		s  []int
		em method
		ee error
	}{
		{[]int{4, 2}, method{"swap", 0, 1}, nil},
		{[]int{1, 5, 4, 3, 2, 6}, method{"reverse", 1, 4}, nil},
	}
	for _, c := range cases {
		s0 := make([]int, len(c.s))
		copy(s0, c.s)
		gm, ge := sortMethod(s0)
		if ge != c.ee || *gm != c.em {
			t.Errorf("%v, Expectd: %s, %s, Got: %s, %s", c.s, &c.em, c.ee, gm, ge)
		}
	}
}
