package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ss := readInput(os.Stdin)
	for _, s := range ss {
		if funny(s) {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
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

func funny(s string) bool {
	l := len(s)
	for i := 0; i <= l/2; i++ {
		if abs(s[i], s[i+1]) != abs(s[l-i-1], s[l-i-2]) {
			return false
		}
	}
	return true
}

func abs(a, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}
