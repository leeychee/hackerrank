package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	g := readInput(os.Stdin)
	maxArea := g.maxGoodArea()
	fmt.Println(maxArea)
}

func readInput(rd io.Reader) grid {
	var r, c int
	fmt.Fscanf(rd, "%d %d\n", &r, &c)
	g := make(grid, r)
	for i := 0; i < r; i++ {
		var s string
		fmt.Fscanf(rd, fmt.Sprintf("%%%ds\n", c), &s)
		g[i] = []byte(s)
	}
	return g
}

type cell struct{ i, j int }
type grid [][]byte

type plus struct {
	center cell
	_type  byte
	d      int
}

func (p *plus) String() string {
	return fmt.Sprintf("([%d, %d]: %d)", p.center.i, p.center.j, p.d)
}

func (p *plus) i() int   { return p.center.i }
func (p *plus) j() int   { return p.center.j }
func (p *plus) len() int { return 4*p.d + 1 }

func (p *plus) overlap(p0 *plus) bool {
	m := make(map[cell]struct{})
	for i := p.i() - p.d; i <= p.i()+p.d; i++ {
		m[cell{i, p.j()}] = struct{}{}
	}
	for j := p.j() - p.d; j <= p.j()+p.d; j++ {
		m[cell{p.i(), j}] = struct{}{}
	}
	for i := p0.i() - p0.d; i <= p0.i()+p0.d; i++ {
		if _, ok := m[cell{i, p0.j()}]; ok {
			return true
		}
	}
	for j := p0.j() - p0.d; j <= p0.j()+p0.d; j++ {
		if _, ok := m[cell{p0.i(), j}]; ok {
			return true
		}
	}
	return false
}

var errUnableExpand = errors.New("unable expand")

func (p *plus) expand(g grid) error {
	d0 := p.d + 1
	i, j := p.center.i, p.center.j
	if !g.valid(i+d0, j) || g[i+d0][j] != p._type {
		return errUnableExpand
	}
	if !g.valid(i-d0, j) || g[i-d0][j] != p._type {
		return errUnableExpand
	}
	if !g.valid(i, j-d0) || g[i][j-d0] != p._type {
		return errUnableExpand
	}
	if !g.valid(i, j+d0) || g[i][j+d0] != p._type {
		return errUnableExpand
	}
	p.d++
	return nil
}

func (g grid) r() int {
	return len(g)
}

func (g grid) c() int {
	return len(g[0])
}

func (g grid) valid(i, j int) bool {
	return i >= 0 && i < g.r() && j >= 0 && j < g.c()
}

func (g grid) maxGoodArea() int {
	var ps []*plus
	var pm = make(map[string]struct{})
	for i := 0; i < g.r(); i++ {
		for j := 0; j < g.c(); j++ {
			if g[i][j] != 'G' {
				continue
			}
			c := cell{i, j}
			p := &plus{center: c, _type: 'G'}
			for k := 1; ; k++ {
				if _, ok := pm[p.String()]; !ok {
					pm[p.String()] = struct{}{}
					// copy the plus object.
					// assignment makes copy.
					p1 := *p
					ps = append(ps, &p1)
				}
				if err := p.expand(g); err != nil {
					break
				}
			}
		}
	}
	// for i := 0; i < len(ps); i++ {
	// 	fmt.Printf("%s\n", ps[i])
	// }
	// fmt.Println("------------------------------")
	max := 0
	for i := 0; i < len(ps)-1; i++ {
		for j := i + 1; j < len(ps); j++ {
			if !ps[i].overlap(ps[j]) {
				a := ps[i].len() * ps[j].len()
				// fmt.Printf("%s %s, %d*%d=%d\n", ps[i], ps[j], ps[i].len(), ps[j].len(), a)
				if a > max {
					max = a
				}
			}
		}
	}
	return max
}
