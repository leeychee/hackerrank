package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ts := readInput(os.Stdin)
	for _, t := range ts {
		fmt.Printf("%d\n", t.feast())
	}
}

type trip struct{ n, c, m int }

func (t *trip) String() string {
	return fmt.Sprintf("n: %d, c: %d, m: %d", t.n, t.c, t.m)
}

func (t *trip) feast() (f int) {
	f += t.n / t.c
	f1 := f / t.m
	y1 := f % t.m
	for f1 > 0 {
		f += f1
		f1, y1 = (f1+y1)/t.m, (f1+y1)%t.m
	}
	return
}

func readInput(r io.Reader) []*trip {
	var t int
	fmt.Fscanf(r, "%d\n", &t)

	s := make([]*trip, t)
	for i := 0; i < t; i++ {
		t0 := trip{}
		fmt.Fscanf(r, "%d %d %d\n", &t0.n, &t0.c, &t0.m)
		s[i] = &t0
	}
	return s
}
