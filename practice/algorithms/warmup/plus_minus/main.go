package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	si := readInput(os.Stdin)
	p, z, n := plusMinus(si)
	l := float32(len(si))
	fmt.Printf("%.6f\n%.6f\n%.6f", float32(p)/l, float32(n)/l, float32(z)/l)
}

func readInput(r io.Reader) []int {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	s := make([]int, n)
	for i := range s {
		fmt.Fscanf(r, "%d", &s[i])
	}
	return s
}

func plusMinus(s []int) (positive, zeros, negative int) {
	for _, i := range s {
		if i > 0 {
			positive++
		} else if i == 0 {
			zeros++
		} else {
			negative++
		}
	}
	return
}
