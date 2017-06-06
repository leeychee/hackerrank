package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	g := readInput(os.Stdin)
	pluses := g.validPluses()
	if len(pluses) >= 2 {
		area := len(pluses[0]) * len(pluses[1])
		fmt.Println(area)
	}
	fmt.Println(0)
}

func readInput(rd io.Reader) grid {
	var r, c int
	fmt.Fscanf(rd, "%d %d\n", &r, &c)
	g := make(grid, r)
	for i := 0; i < r; i++ {
		var s string
		fmt.Fscanf(rd, fmt.Sprintf("%%%ds", c), &s)
		g[i] = []byte(s)
	}
	return g
}

type grid [][]byte
type plus []cell
type cell struct{ i, j int }

func (g grid) validPluses() []plus {
	return nil
}
