package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s, xs := readInput(os.Stdin)
	fmt.Print(toString(s, xs))
}

func toString(s string, xs []int) string {
	m := uniformWeightMap(s)
	var buf bytes.Buffer
	for _, x := range xs {
		if _, ok := m[x]; ok {
			fmt.Fprintln(&buf, "Yes")
		} else {
			fmt.Fprintln(&buf, "No")
		}
	}
	return buf.String()
}

func readInput(r io.Reader) (s string, xs []int) {
	fmt.Fscanln(r, &s)
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	xs = make([]int, n)
	for i := range xs {
		fmt.Fscanf(r, "%d\n", &xs[i])
	}
	return
}

func uniformWeightMap(s string) map[int]struct{} {
	s = strings.ToLower(s)
	var cur byte
	var count int
	m := make(map[int]struct{})
	for _, b := range []byte(s) {
		if cur == byte(0) || b == cur {
			count++
		} else {
			count = 1
		}
		cur = b
		c := weight(cur) * count
		m[c] = struct{}{}
	}
	return m
}

func weight(b byte) int {
	return int(b-'a') + 1
}
