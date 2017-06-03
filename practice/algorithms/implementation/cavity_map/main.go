package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	m := readInput(os.Stdin)
	fmt.Printf("%s", m)
}

type matrix [][]int

func newMatrix(n int) matrix {
	m := make(matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	return m
}

func (m matrix) isCavity(i, j int) bool {
	if i == 0 || j == 0 || i == len(m)-1 || j == len(m)-1 {
		return false
	}
	if m[i][j] <= m[i-1][j] || m[i][j] <= m[i+1][j] || m[i][j] <= m[i][j-1] || m[i][j] <= m[i][j+1] {
		return false
	}
	return true
}

func (m matrix) String() string {
	var buf bytes.Buffer
	n := len(m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m.isCavity(i, j) {
				fmt.Fprint(&buf, "X")
			} else {
				fmt.Fprintf(&buf, "%d", m[i][j])
			}
		}
		if i < n-1 {
			fmt.Fprint(&buf, "\n")
		}
	}
	return buf.String()
}

func readInput(r io.Reader) matrix {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	m := newMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(r, "%1d", &m[i][j])
		}
		fmt.Fscanln(r)
	}
	return m
}
