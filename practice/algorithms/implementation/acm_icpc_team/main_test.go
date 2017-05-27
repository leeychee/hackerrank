package main

import "testing"

var cases = []struct {
	t               [][]int
	etopics, eteams int
}{
	{[][]int{
		{1, 0, 1, 0, 1},
		{1, 1, 1, 0, 0},
		{1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1},
	}, 5, 2},
}

func TestMaxTopicsOfTwoPersons(t *testing.T) {
	for _, c := range cases {
		gtopics, gteams := maxTopicsOfTwoPersons(c.t)
		if gtopics != c.etopics || gteams != c.eteams {
			t.Errorf("Topics Expected: %d, Got: %d\n", c.etopics, gtopics)
			t.Errorf("Teams Expected: %d, Got: %d\n", c.eteams, gteams)
		}
	}
}

func TestOr(t *testing.T) {
	cases := []struct {
		t1, t2 []int
		e      int
	}{
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 0},
		{[]int{1, 1, 0, 1, 1}, []int{0, 0, 1, 0, 0}, 5},
		{[]int{1, 0, 0, 1, 1}, []int{0, 0, 1, 0, 0}, 4},
	}
	for _, c := range cases {
		g, _ := or(c.t1, c.t2)
		if g != c.e {
			t.Errorf("Expected: %d, Got: %d\n", c.e, g)
		}
	}
}

func BenchmarkMaxTopicsOfTwoPersons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxTopicsOfTwoPersons(cases[0].t)
	}
}
