package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	l, ss := readInput(os.Stdin)
	// fmt.Println(l)
	// fmt.Println(ss)
	for _, s := range ss {
		c := pctype(l, s)
		fmt.Println(c)
	}
}

type lane []int
type service struct{ entry, exit int }

func readInput(r io.Reader) ([]int, []*service) {
	var n, t int
	fmt.Fscanf(r, "%d %d\n", &n, &t)
	l := make(lane, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &l[i])
	}
	//fmt.Fscanln(r)
	ss := make([]*service, t)
	for i := 0; i < t; i++ {
		s := &service{}
		fmt.Fscanf(r, "%d %d\n", &s.entry, &s.exit)
		ss[i] = s
	}
	return l, ss
}

func pctype(l lane, s *service) (c int) {
	c = 3
	for i := s.entry; i <= s.exit && i < len(l); i++ {
		if l[i] < c {
			c = l[i]
		}
	}
	return
}
