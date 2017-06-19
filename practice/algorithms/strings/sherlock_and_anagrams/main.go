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
		fmt.Println(anagramCount(s))
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

func anagramCount(s string) int {
	subs := make(map[string]int)
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			bs := []byte(s)[i:j]
			sort.Sort(bytesSorter(bs))
			subs[string(bs)]++
		}
	}
	r := 0
	for _, count := range subs {
		r += count * (count - 1) / 2
	}
	return r
}

type bytesSorter []byte

func (bs bytesSorter) Len() int {
	return len(bs)
}
func (bs bytesSorter) Less(i, j int) bool {
	return bs[i] < bs[j]
}
func (bs bytesSorter) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}
