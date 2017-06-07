package main

import "testing"
import "reflect"

func TestRotate(t *testing.T) {
	var cases = []struct {
		a []int
		e []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 1}},
	}
	for _, c := range cases {
		a0 := make([]int, len(c.a))
		copy(a0, c.a)
		rotate(a0)
		if !reflect.DeepEqual(a0, c.e) {
			t.Errorf("rotate(%v), Expected: %v, Got: %v", c.a, c.e, a0)
		}
	}
}

func TestRotateSort(t *testing.T) {
	var cases = []struct {
		a []int
		e error
	}{
		{[]int{3, 2, 1}, errUnableSort},
		{[]int{3, 1, 2}, nil},
		{[]int{1, 3, 4, 2}, nil},
		{[]int{1, 2, 3, 5, 4}, errUnableSort},
		{[]int{4, 1, 3, 2}, nil},
	}
	for _, c := range cases {
		a0 := make([]int, len(c.a))
		copy(a0, c.a)
		g := rotateSort(a0)
		if g != c.e {
			t.Errorf("%v, %v, Expected: %s, Got: %s", c.a, a0, c.e, g)
		}
	}
}

func BenchmarkRotateSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateSort([]int{4, 1, 3, 2, 7, 5, 8})
	}
}
