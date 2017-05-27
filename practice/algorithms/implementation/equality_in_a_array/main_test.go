package main

import "testing"

var cases = []struct {
	s []int
	e int
}{
	{[]int{3, 3, 2, 1, 3}, 2},
}

func TestMakeItEqual(t *testing.T) {
	for _, c := range cases {
		_, g := makeItEqual(c.s)
		if g != c.e {
			t.Errorf("Expected: %d, Got: %d\n", c.e, g)
		}
	}
}

func BenchmarkMakeItEqaul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeItEqual(cases[0].s)
	}
}
