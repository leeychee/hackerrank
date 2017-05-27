package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	n, q, o := readInput(os.Stdin)
	a := attackCells(n, q, o)
	fmt.Printf("%d\n", len(a))
}

type cell struct{ r, c int }

func (c *cell) String() string {
	return fmt.Sprintf("[%d, %d]", c.r, c.c)
}

func (c *cell) same(d *cell) bool {
	return c.c == d.c && c.r == d.r
}

func (c *cell) attacked(c2 *cell) bool {
	return !c.same(c2) && (c.r == c2.r ||
		c.c == c2.c ||
		(c.c-c2.c) == (c.r-c2.r) ||
		(c.c-c2.c) == (c2.r-c.r))
}

func (c *cell) inCells(cs []cell) bool {
	for _, v := range cs {
		if v.same(c) {
			return true
		}
	}
	return false
}

func readInput(r io.Reader) (int, cell, []cell) {
	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)
	var queen cell
	fmt.Fscanf(r, "%d %d\n", &queen.r, &queen.c)

	obstacles := make([]cell, k)
	for i := 0; i < k; i++ {
		fmt.Fscanf(r, "%d %d\n", &obstacles[i].r, &obstacles[i].c)
	}
	return n, queen, obstacles
}

func readOutput(r io.Reader) (s int) {
	fmt.Fscanf(r, "%d", &s)
	return
}

func attackCells(n int, q cell, s []cell) []cell {
	a := make([]cell, 0)
	cmap := make(map[cell]struct{})
	for _, v := range s {
		cmap[v] = struct{}{}
	}
	a = attackCellsInRow(n, q, cmap, a, 1)
	a = attackCellsInRow(n, q, cmap, a, -1)
	a = attackCellsInCol(n, q, cmap, a, 1)
	a = attackCellsInCol(n, q, cmap, a, -1)
	a = attackCellsInDiagonal(n, q, cmap, a, -1, -1)
	a = attackCellsInDiagonal(n, q, cmap, a, -1, 1)
	a = attackCellsInDiagonal(n, q, cmap, a, 1, 1)
	a = attackCellsInDiagonal(n, q, cmap, a, 1, -1)
	return a
}

func attackCellsInRow(n int, q cell, s map[cell]struct{}, a []cell, stepr int) (r []cell) {
	r = a
	for i := q.r + stepr; i <= n && i > 0; i += stepr {
		c := cell{i, q.c}
		if _, ok := s[c]; !ok {
			r = append(r, c)
		} else {
			break
		}
	}
	return
}

func attackCellsInCol(n int, q cell, s map[cell]struct{}, a []cell, stepc int) (r []cell) {
	r = a
	for j := q.c + stepc; j <= n && j > 0; j += stepc {
		c := cell{q.r, j}
		if _, ok := s[c]; !ok {
			r = append(r, c)
		} else {
			break
		}
	}
	return
}

func attackCellsInDiagonal(n int, q cell, s map[cell]struct{}, a []cell, stepr, stepc int) (r []cell) {
	r = a
	i, j := q.r, q.c
	for {
		i, j = i+stepr, j+stepc
		if i > n || i < 1 || j > n || j < 1 {
			break
		}
		c := cell{i, j}
		if _, ok := s[c]; !ok {
			r = append(r, c)
		} else {
			break
		}
	}
	return
}
