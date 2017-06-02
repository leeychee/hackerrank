package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	matrixes := readInput(os.Stdin)
	for _, m := range matrixes {
		possible := m.canSwapToPure()
		if possible {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}

type matrix [][]int

func newMatrix(n int) matrix {
	m := make(matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	return m
}

func (m matrix) String() string {
	n := len(m)
	var buf bytes.Buffer
	for i, l := range m {
		for j, c := range l {
			fmt.Fprintf(&buf, "%d", c)
			if j < n-1 {
				fmt.Fprintf(&buf, ",")
			}
		}
		if i < n-1 {
			fmt.Fprintln(&buf)
		}
	}
	return buf.String()
}

func (m matrix) canSwapToPure() bool {
	for i := 0; i < len(m); i++ {
		tiC := 0
		for j := 0; j < len(m); j++ {
			tiC += m[j][i]
		}
		if tiC != len(m) {
			return false
		}
	}
	return true
}

func readInput(r io.Reader) []matrix {
	var q int
	fmt.Fscanf(r, "%d\n", &q)
	ms := make([]matrix, q)
	for i := 0; i < q; i++ {
		var n int
		fmt.Fscanf(r, "%d\n", &n)
		m := newMatrix(n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				fmt.Fscanf(r, "%d", &m[j][k])
			}
		}
		ms[i] = m
	}
	return ms
}
