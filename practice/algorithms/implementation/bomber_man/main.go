package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	g, n := readInput(os.Stdin)
	g.after(n)
	fmt.Print(g)
}

func readInput(rd io.Reader) (*grid, int) {
	var r, c, n int
	fmt.Fscanf(rd, "%d %d %d\n", &r, &c, &n)
	g := newGrid(r, c)
	g.read(rd)
	return g, n
}

type grid struct {
	r, c int
	g    [][]bomb
}

func newGrid(r, c int) *grid {
	g := &grid{r: r, c: c, g: make([][]bomb, r)}
	for i := 0; i < r; i++ {
		g.g[i] = make([]bomb, c)
	}
	return g
}

func (g *grid) deepCopy() *grid {
	g0 := newGrid(g.r, g.c)
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			b := g.get(i, j)
			if b != emptyBomb {
				g0.g[i][j] = b
			}
		}
	}
	return g0
}

func (g *grid) read(r io.Reader) {
	for i := 0; i < g.r; i++ {
		var s string
		fmt.Fscanf(r, fmt.Sprintf("%%%ds\n", g.c), &s)
		for j, b := range []byte(s) {
			if b == 'O' {
				g.g[i][j] = newBomb()
			}
		}
	}
}

func (g *grid) String() string {
	var buf bytes.Buffer
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			if g.g[i][j] == emptyBomb {
				fmt.Fprintf(&buf, "%s", ".")
			} else {
				fmt.Fprintf(&buf, "%s", "O")
			}
		}
		if i < g.r-1 {
			fmt.Fprintln(&buf)
		}
	}
	return buf.String()
}

func (g *grid) bomb(i, j int) {
	if i >= 0 && i < g.r && j >= 0 && j < g.c {
		g.g[i][j] = newBomb()
	}
}

func (g *grid) empty(i, j int) {
	if i >= 0 && i < g.r && j >= 0 && j < g.c {
		g.g[i][j] = emptyBomb
	}
}

func (g *grid) get(i, j int) bomb {
	if i >= 0 && i < g.r && j >= 0 && j < g.c {
		return g.g[i][j]
	}
	return emptyBomb
}

func (g *grid) tick() []*cell {
	explosion := make([]*cell, 0)
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			b := g.get(i, j)
			if b == emptyBomb {
				continue
			}
			if err := (&(g.g[i][j])).tick(); err != nil {
				explosion = append(explosion, &cell{i, j})
			}
		}
	}
	return explosion
}

func (g *grid) after(n int) {
	gs := make([]*grid, n+1)
	c := 1
	for ; c <= n; c++ {
		explosion := g.tick()
		if c%2 != 0 {
			// explode
			for _, e := range explosion {
				g.empty(e.i, e.j)
				for _, k := range []int{e.i - 1, e.i + 1} {
					if !g.get(k, e.j).bombed() {
						g.empty(k, e.j)
					}
				}
				for _, k := range []int{e.j - 1, e.j + 1} {
					if !g.get(e.i, k).bombed() {
						g.empty(e.i, k)
					}
				}
			}
		} else {
			g.fill()
		}
		gs[c] = g.deepCopy()
		if c > 3 && gs[3].deepEqual(g) {
			break
		}
		// fmt.Printf("--------------%d-----------------%t\n", c, gs[3] != nil)
		// fmt.Println(g)
		// fmt.Println("------------------------------")
	}
	if n > c {
		i := (n-3)%(c-3) + 3
		// fmt.Printf("n: %d, c: %d, i: %d\n", n, c, i)
		g.g = gs[i].g
	}
	return
}

func (g *grid) hasBomb(i, j int) bool {
	return g.get(i, j) != emptyBomb
}

func (g *grid) fill() {
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			if g.get(i, j) == emptyBomb {
				g.bomb(i, j)
			}
		}
	}
}

func (g *grid) equal(g0 *grid) bool {
	if g.r != g0.r || g.c != g0.c {
		return false
	}
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			if g.hasBomb(i, j) != g0.hasBomb(i, j) {
				return false
			}
		}
	}
	return true
}

func (g *grid) deepEqual(g0 *grid) bool {
	if g.r != g0.r || g.c != g0.c {
		return false
	}
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			b := g.get(i, j)
			b0 := g0.get(i, j)
			if b != b0 {
				return false
			}
		}
	}
	return true
}

type cell struct{ i, j int }

var errBombed = errors.New("bombed")
var emptyBomb bomb

type bomb int

func newBomb() bomb {
	return bomb(3)
}

func (b *bomb) tick() error {
	*b--
	if *b < 1 {
		*b = -1
		return errBombed
	}
	return nil
}

func (b bomb) bombed() bool {
	return b == -1
}
