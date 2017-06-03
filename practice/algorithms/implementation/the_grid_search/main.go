package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	cases := readInput(os.Stdin)
	for _, v := range cases {
		//fmt.Println(v.g)
		//fmt.Println(v.p)
		if _, _, err := v.g.contain(v.p); err == nil {
			//fmt.Printf("%d, %d\n", x, y)
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type _case struct{ g, p grid }

type grid [][]int

func (g grid) w() int {
	return len(g[0])
}

func (g grid) h() int {
	return len(g)
}

// ErrNotContain not contain
var ErrNotContain = errors.New("not contain")

func (g grid) contain(p grid) (int, int, error) {
	if g.w() < p.w() || g.h() < p.h() {
		return 0, 0, ErrNotContain
	}
	for i := 0; i <= g.h()-p.h(); i++ {
		for j := 0; j <= g.w()-p.w(); j++ {
			ctd := true
		once:
			for k := 0; k < p.h(); k++ {
				for l := 0; l < p.w(); l++ {
					if g[i+k][j+l] != p[k][l] {
						ctd = false
						break once
					}
				}
			}
			if ctd {
				return i, j, nil
			}
		}
	}
	return 0, 0, ErrNotContain
}

func readInput(r io.Reader) []*_case {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	gs := make([]*_case, n)
	for i := 0; i < n; i++ {
		gs[i] = readCase(r)
	}
	return gs
}

func readCase(r io.Reader) *_case {
	g := readGrid(r)
	p := readGrid(r)
	return &_case{g, p}
}

func readGrid(reader io.Reader) grid {
	var r, c int
	fmt.Fscanf(reader, "%d %d\n", &r, &c)
	g := make(grid, r)
	for i := 0; i < r; i++ {
		c0 := make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Fscanf(reader, "%1d", &c0[j])
		}
		fmt.Fscanln(reader)
		g[i] = c0
	}
	return g
}
