package main

import "testing"
import "bytes"

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
