package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ss := readInput(os.Stdin)
	for _, s := range ss {
		if includeHackerrank(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func readInput(r io.Reader) []string {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Fscanf(r, "%s\n", &ss[i])
	}
	return ss
}

func includeHackerrank(s string) bool {
	return include(s, "hackerrank")
}

func include(s, s0 string) bool {
	bs0 := []byte(s0)
	cursor := 0
	for _, b := range []byte(s) {
		if cursor < len(bs0) {
			if bs0[cursor] == b {
				cursor++
			}
		}
	}
	return cursor == len(bs0)
}
