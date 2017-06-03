package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, k, ts := readInput(os.Stdin)
	fmt.Println(specialNum(k, ts))
}

func readInput(r io.Reader) (n, k int, ts []int) {
	fmt.Fscanf(r, "%d %d\n", &n, &k)
	ts = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &ts[i])
	}
	return
}

func specialNum(k int, ts []int) (s int) {
	p := 1
	for _, t := range ts {
		for j := 1; j <= t; j++ {
			if j == (j-1)/k+p {
				s++
			}
		}
		p += (t-1)/k + 1
	}
	return
}
