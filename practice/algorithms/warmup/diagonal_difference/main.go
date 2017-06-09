package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var m = make([][]int, n)
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

	fmt.Printf("%d", distance(d0, d1))
}

func distance(a, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d
}
