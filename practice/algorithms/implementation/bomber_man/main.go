package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	g, n := readInput(os.Stdin)
	fmt.Printf("%d %d %d\n", g.r, g.c, n)
	fmt.Println(g)
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
	g    [][]*bomb
}

func newGrid(r, c int) *grid {
	g := &grid{r: r, c: c, g: make([][]*bomb, r)}
	for i := 0; i < r; i++ {
		g.g[i] = make([]*bomb, c)
	}
	return g
}

// T time
var T = 3

func (g *grid) read(r io.Reader) {
	for i := 0; i < g.r; i++ {
		var s string
		fmt.Fscanf(r, fmt.Sprintf("%%%ds\n", g.c), &s)
		for j, b := range []byte(s) {
			if b == 'O' {
				g.g[i][j] = &bomb{T}
			}
		}
	}
}

func (g *grid) String() string {
	var buf bytes.Buffer
	for i := 0; i < g.r; i++ {
		for j := 0; j < g.c; j++ {
			if g.g[i][j] == nil {
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

func (g *grid) after(n int) {
	//c :=
}

type bomb struct{ t int }
