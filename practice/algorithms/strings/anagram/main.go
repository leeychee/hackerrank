package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	ss := readStringsN(os.Stdin)
	for _, s := range ss {
		fmt.Println(makeAnagram(s))
	}
}

func readStringsN(r io.Reader) []string {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	ss := make([]string, t)
	for i := range ss {
		fmt.Fscanf(r, "%s\n", &ss[i])
	}
	return ss
}

type bytesSorter []byte

func (b bytesSorter) Len() int {
	return len(b)
}

func (b bytesSorter) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b bytesSorter) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func isAnagram(s1, s2 string) bool {
	bs1 := []byte(s1)
	bs2 := []byte(s2)
	sort.Sort(bytesSorter(bs1))
	sort.Sort(bytesSorter(bs2))
	if string(bs1) == string(bs2) {
		return true
	}
	return false
}

func makeAnagram(s string) int {
	n := len(s)
	if n%2 != 0 {
		return -1
	}
	bs := []byte(s)
	m1 := toMap(bs[:n/2])
	m2 := toMap(bs[n/2:])
	var r0, r1 int
	for b, v := range m1 {
		if m2[b]-v > 0 {
			r0 += m2[b] - v
		} else if m2[b]-v < 0 {
			r1 -= m2[b] - v
		}
	}
	if r1 > r0 {
		return r1
	}
	return r0
}

func toMap(bs []byte) map[byte]int {
	m := make(map[byte]int)
	for _, b := range bs {
		m[b]++
	}
	return m
}
