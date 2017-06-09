package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ss := readInput(os.Stdin)
	fmt.Printf("%d\n", gem(ss))
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

func gem(ss []string) int {
	var gs []byte
	for i, s := range ss {
		bs := []byte(s)
		if i == 0 {
			gs = bs
		} else {
			gs = intersection(gs, bs)
		}
	}
	return len(gs)
}

func intersection(s0, s1 []byte) []byte {
	m0 := make(map[byte]bool)
	for _, b := range s0 {
		m0[b] = true
	}
	m := make(map[byte]bool)
	var s []byte
	for _, b := range s1 {
		if !m[b] && m0[b] {
			m[b] = true
			s = append(s, b)
		}
	}
	return s
}
