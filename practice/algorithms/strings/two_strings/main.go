package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ps := readPairs(os.Stdin)
	for _, p := range ps {
		if haveCommonSubstring(p.s1, p.s2) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type pair struct{ s1, s2 string }

func readPairs(r io.Reader) []pair {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	ps := make([]pair, t)
	for i := 0; i < t; i++ {
		ps[i] = pair{}
		fmt.Fscanf(r, "%s\n", &ps[i].s1)
		fmt.Fscanf(r, "%s\n", &ps[i].s2)
	}
	return ps
}

func haveCommonSubstring(s1, s2 string) bool {
	m1 := make(map[byte]bool)
	for _, b := range []byte(s1) {
		m1[b] = true
	}
	for _, b := range []byte(s2) {
		if m1[b] {
			return true
		}
	}
	return false
}
