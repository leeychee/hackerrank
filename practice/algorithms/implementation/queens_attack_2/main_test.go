package main

import "testing"
import "os"

func TestAttackCells(t *testing.T) {
	var cases = []struct {
		n         int
		queen     cell
		obstacles []cell
		e         int
	}{
		{4, cell{4, 4}, []cell{}, 9},
		{5, cell{4, 3}, []cell{{5, 5}, {4, 2}, {2, 3}}, 10},
	}
	for _, c := range cases {
		g := attackCells(c.n, c.queen, c.obstacles)
		if len(g) != c.e {
			t.Logf("%v\n", g)
			t.Errorf("Expected: %d, Got: %d\n", c.e, len(g))
		}
	}
}

func TestAttackCellsFromFile(t *testing.T) {
	f, _ := os.Open("./input13.txt")
	n, q, o := readInput(f)
	g := attackCells(n, q, o)
	f, _ = os.Open("./output13.txt")
	e := readOutput(f)
	if len(g) != e {
		t.Errorf("Expected: %d, Got: %d\n", e, len(g))
	}
}

func TestSame(t *testing.T) {
	var cases = []struct {
		c1, c2 cell
		e      bool
	}{
		{cell{4, 4}, cell{1, 1}, false},
	}
	for _, c := range cases {
		g := c.c1.same(&c.c2)
		if g != c.e {
			t.Errorf("%s <-> %s, Expected: %t, Got: %t\n", &c.c1, &c.c2, c.e, g)
		}
	}
}

func TestAttacked(t *testing.T) {
	var cases = []struct {
		c1, c2 cell
		e      bool
	}{
		{cell{4, 4}, cell{4, 4}, false},
		{cell{4, 4}, cell{1, 1}, true},
		{cell{4, 4}, cell{1, 2}, false},
		{cell{4, 4}, cell{2, 1}, false},
	}
	for _, c := range cases {
		g := c.c1.attacked(&c.c2)
		if g != c.e {
			t.Errorf("%s <-> %s, Expected: %t, Got: %t\n", &c.c1, &c.c2, c.e, g)
		}
	}

}

func BenchmarkAttackCells(b *testing.B) {
	for i := 0; i < b.N; i++ {
		attackCells(4, cell{4, 4}, []cell{})
	}
}

func BenchmarkAttackCellsFromFile(b *testing.B) {
	f, _ := os.Open("./input13.txt")
	n, q, o := readInput(f)
	for i := 0; i < b.N; i++ {
		attackCells(n, q, o)
	}
}
