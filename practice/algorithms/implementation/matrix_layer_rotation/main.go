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
	fmt.Println(m)
	fmt.Println(r)
}

func readInput(r io.Reader) (matrix, int) {
	var m, n, rotation int
	fmt.Fscanf(r, "%d %d %d\n", &m, &n, &rotation)

	mtx := make(matrix, m)
	for i := range mtx {
		mtx[i] = make([]int, n)
		for j := range mtx[i] {
			fmt.Fscanf(r, "%d", &mtx[i][j])
		}
		// fmt.Fscanln(r)
	}
	return mtx, rotation
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

func (m matrix) rotate() error {
	r, c := len(m), len(m[0])
	if r%2 != 0 && c%2 != 0 {
		return errUnableRotate
	}

	return nil
}
