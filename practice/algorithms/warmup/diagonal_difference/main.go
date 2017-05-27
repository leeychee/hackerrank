package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var m [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &m[i][j])
		}
	}

	var d0, d1 int
	for i := 0; i < n; i++ {
		d0 += m[i][i]
	}
	for i := 0; i < n; i++ {
		d1 += m[n-i-1][i]
	}

	fmt.Printf("%.0f", math.Abs(float64(d0-d1)))
}
