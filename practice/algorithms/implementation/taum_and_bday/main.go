package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	plans := readInput(os.Stdin)
	for _, p := range plans {
		fmt.Println(p.spend())
	}
}

type plan struct{ b, w, x, y, z int }

func (p *plan) String() string {
	return fmt.Sprintf("[%d, %d] [%d, %d, %d]", p.b, p.w, p.x, p.y, p.z)
}

func (p *plan) spend() int {
	xz := p.y + p.z
	if xz > p.x {
		xz = p.x
	}
	yz := p.x + p.z
	if yz > p.y {
		yz = p.y
	}
	return p.b*xz + p.w*yz
}

func readInput(r io.Reader) []*plan {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	plans := make([]*plan, t)
	for i := 0; i < t; i++ {
		p := plan{}
		fmt.Fscanf(r, "%d %d\n", &p.b, &p.w)
		fmt.Fscanf(r, "%d %d %d\n", &p.x, &p.y, &p.z)
		plans[i] = &p
	}
	return plans
}
