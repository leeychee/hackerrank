package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s, n := readInput(os.Stdin)
	c := countRepeatedString('a', s, n)
	fmt.Printf("%d\n", c)
}

func readInput(r io.Reader) (string, int) {
	var s string
	fmt.Fscanf(r, "%s\n", &s)
	var n int
	fmt.Fscanf(r, "%d", &n)
	return s, n
}

func countRepeatedString(a byte, s string, n int) int {
	cis := 0
	cil := 0
	for i, c := range []byte(s) {
		if c == a {
			cis++
			if i < n%len(s) {
				cil++
			}
		}
	}
	return n/len(s)*cis + cil
}
