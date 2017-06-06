package main

import "testing"
import "bytes"
import "os"

func TestAfter(t *testing.T) {
	var cases = []struct {
		r, c, n  int
		ogs, fgs string
	}{
		{6, 7, 3,
			`.......
...O...
....O..
.......
OO.....
OO.....`, `OOO.OOO
OO...OO
OOO...O
..OO.OO
...OOOO
...OOOO`},
	}

	for _, c := range cases {
		g0 := newGrid(c.r, c.c)
		g0.read(bytes.NewReader([]byte(c.ogs)))
		g1 := newGrid(c.r, c.c)
		g1.read(bytes.NewReader([]byte(c.fgs)))
		g0.after(c.n)
		if g0.deepEqual(g1) {
			t.Errorf("%s\nAfter %d, Expected:\n%s, Got:\n%s", c.ogs, c.n, g1, g0)
		}
	}
}

func TestAfterFromFile(t *testing.T) {
	var cases = []struct {
		fis, fos string
	}{
		{"input13.txt", "output13.txt"},
		{"input24.txt", "output24.txt"},
	}
	for _, c := range cases {
		fi, _ := os.Open(c.fis)
		g, n := readInput(fi)
		g.after(n)
		e := newGrid(g.r, g.c)
		fo, _ := os.Open(c.fos)
		e.read(fo)
		if !g.equal(e) {
			t.Errorf("Test case %s failed.", c.fis)
		}
	}
}

func BenchmarkAfter(b *testing.B) {
	fin, _ := os.Open("./input24.txt")
	g, n := readInput(fin)
	g0 := g.deepCopy()
	for i := 0; i < b.N; i++ {
		g.after(n)
		g = g0.deepCopy()
	}
}
