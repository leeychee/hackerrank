package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ss := readStringsN(os.Stdin)
	for _, s := range ss {
		fmt.Println(construct(s))
	}
}

func readStringsN(r io.Reader) []string {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	ss := make([]string, t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%s\n", &ss[i])
	}
	return ss
}

// In fact, we don't need to a complex alg.
// For there is not limit for substrings, we can just use substing
// which len(sub) == 1. Then we just need to count uniq byte.
func construct(s string) (step int) {
	bs := []byte(s)
	bm := make(map[byte]int)
	for _, b := range bs {
		if bm[b] == 0 {
			step++
		}
		bm[b]++
	}
	return
}
