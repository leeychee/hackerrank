package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ss := readInput(os.Stdin)
	for _, s := range ss {
		s0 := alternate(s)
		fmt.Println(len(s) - len(s0))
	}
}

func readInput(r io.Reader) []string {
	var n int
	fmt.Fscanf(r, "%d", &n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Fscanf(r, "%s\n", &ss[i])
	}
	return ss
}

func alternate(s string) string {
	var c byte
	var s0 []byte
	for _, b := range []byte(s) {
		if c == byte(0) || b != c {
			s0 = append(s0, b)
			c = b
		}
	}
	return string(s0)
}
