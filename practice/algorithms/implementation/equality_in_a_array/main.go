package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := readInput(os.Stdin)
	_, min := makeItEqual(s)
	fmt.Printf("%d\n", min)
}

func readInput(r io.Reader) []int {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &s[i])
	}
	return s
}

func makeItEqual(s []int) (int, int) {
	vmap := make(map[int]int)
	for _, v := range s {
		vmap[v]++
	}
	maxv := 0
	maxc := 0
	for v, c := range vmap {
		if maxc < c {
			maxc = c
			maxv = v
		}
	}
	return maxv, len(s) - maxc
}
