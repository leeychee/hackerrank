package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	ts := readInput(os.Stdin)
	for _, t := range ts {
		// fmt.Println(t)
		if err := rotateSort(t); err == nil {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func readInput(r io.Reader) [][]int {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	ts := make([][]int, t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanf(r, "%d\n", &n)
		s := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(r, "%d", &s[j])
		}
		ts[i] = s
	}
	return ts
}

var errUnableSort = errors.New("unable sort")

func rotate(a []int) {
	if len(a) == 3 {
		a[0], a[1], a[2] = a[1], a[2], a[0]
	}
}

func rotateMinFirst(a []int) {
	m := 0
	for i := 1; i < 3; i++ {
		if a[i] < a[m] {
			m = i
		}
	}
	for i := 0; i < m; i++ {
		rotate(a)
	}
}

func rotateSort(a []int) error {
	for i := 0; i < len(a)-2; i++ {
		for j := len(a) - 3; j >= i; j-- {
			rotateMinFirst(a[j : j+3])
		}
	}
	if a[len(a)-2] > a[len(a)-1] {
		return errUnableSort
	}
	return nil
}
