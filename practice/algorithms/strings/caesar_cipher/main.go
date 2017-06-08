package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s, k := readInput(os.Stdin)
	fmt.Println(cipher(s, k))
}

func readInput(r io.Reader) (s string, k int) {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fscanf(r, fmt.Sprintf("%%%ds\n", n), &s)
	fmt.Fscanf(r, "%d\n", &k)
	return
}

func cipher(s string, k int) string {
	k = k % 26
	bs := []byte(s)
	for i, b := range bs {
		bs[i] = cipherB(b, k)
	}
	return string(bs)
}

func cipherB(b byte, k int) byte {
	r := b
	if b >= 'a' && b <= 'z' {
		r = b + byte(k)
		if r > 'z' {
			r = 'a' + r - 'z' - 1
		}
	}
	if b >= 'A' && b <= 'Z' {
		r = b + byte(k)
		if r > 'Z' {
			r = 'A' + r - 'Z' - 1
		}
	}
	return r
}
