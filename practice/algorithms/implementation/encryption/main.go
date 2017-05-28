package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	s := readInput(os.Stdin)
	fmt.Println(encryption(s))
}

func readInput(r io.Reader) string {
	var s string
	fmt.Fscanf(r, "%s\n", &s)
	return s
}

func matrixInfo(l int) (row, col int) {
	sqrtL := math.Sqrt(float64(l))
	row = int(math.Floor(sqrtL))
	col = int(math.Ceil(sqrtL))
	if row*col < l {
		row++
	}
	return
}

func encryption(s string) string {
	row, col := matrixInfo(len(s))
	matrix := make([][]byte, row)
	cursor := 0
	for i := 0; i < row; i++ {
		matrix[i] = make([]byte, col)
		for j := 0; j < col; j++ {
			if cursor < len(s) {
				matrix[i][j] = []byte(s)[cursor]
				cursor++
			}
		}
	}
	fmt.Printf("%s", matrix)
	var buf bytes.Buffer
	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			if matrix[i][j] != byte(0) {
				buf.WriteByte(matrix[i][j])
			}
		}
		if j < col-1 {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}
