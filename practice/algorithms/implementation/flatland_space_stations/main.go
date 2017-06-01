package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	n, cs := readInput(os.Stdin)
	d := maxDistance(n, cs)
	fmt.Println(d)
}

func readInput(r io.Reader) (int, []int) {
	var n, m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	cs := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(r, "%d", &cs[i])
	}
	return n, cs
}

func maxDistance(n int, cs []int) (d int) {
	sort.Ints(cs)
	d = cs[0]
	for i := 0; i < len(cs)-1; i++ {
		d0 := (cs[i+1] - cs[i]) / 2
		if d0 > d {
			d = d0
		}
	}
	d1 := n - cs[len(cs)-1] - 1
	if d1 > d {
		d = d1
	}
	return
}
