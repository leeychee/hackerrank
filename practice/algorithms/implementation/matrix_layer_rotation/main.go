package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	m, r := readInput(os.Stdin)
	// fmt.Println(r)
	// fmt.Println(m)
	if err := m.rotate(r); err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

func readInput(r io.Reader) (matrix, int) {
	var m, n, rotation int
	fmt.Fscanf(r, "%d %d %d", &m, &n, &rotation)
	mtx := readMatrix(r, m, n)
	return mtx, rotation
}

func readMatrix(r io.Reader, m, n int) matrix {
	mtx := make(matrix, m)
	for i := range mtx {
		mtx[i] = make([]int, n)
		for j := range mtx[i] {
			p := "%d"
			if j == len(mtx[i])-1 {
				p = "%d\n"
			}
			fmt.Fscanf(r, p, &mtx[i][j])
		}
		// fmt.Fscanln(r)
	}
	return mtx
}

type matrix [][]int

func (m matrix) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Fprint(&buf, m[i][j])
			if j < len(m[i])-1 {
				fmt.Fprint(&buf, " ")
			}
		}
		if i < len(m)-1 {
			fmt.Fprintln(&buf)
		}
	}
	return buf.String()
}

var errUnableRotate = errors.New("unable rotate")

func (m matrix) rotate(times int) error {
	r, c := len(m), len(m[0])
	d := min(r, c)
	if d%2 != 0 {
		return errUnableRotate
	}
	for k := 0; k < d/2; k++ {
		for l := 0; l < times%((r+c-4*k)*2-4); l++ {
			lefttop := m[k][k]
			for j := k; j < c-k-1; j++ {
				m[k][j] = m[k][j+1]
			}
			for i := k; i < r-k-1; i++ {
				m[i][c-k-1] = m[i+1][c-k-1]
			}
			for j := c - k - 1; j > k; j-- {
				m[r-k-1][j] = m[r-k-1][j-1]
			}
			for i := r - k - 1; i > k; i-- {
				m[i][k] = m[i-1][k]
			}
			m[k+1][k] = lefttop
		}
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
