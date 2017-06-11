package main

import (
	"fmt"
	"io"
	"os"

	"strings"
)

func main() {
	ss := readNStrings(os.Stdin)
	for _, s := range ss {
		fmt.Printf("%d\n", makePalindrome(s))
	}
}

func readNStrings(r io.Reader) []string {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	s := make([]string, n)
	for i := range s {
		fmt.Fscanf(r, "%s\n", &s[i])
	}
	return s
}

func makePalindrome(s string) int {
	s = strings.ToLower(s)
	c := 0
	for i := 0; i < len(s)/2; i++ {
		c += distance(s[i], s[len(s)-1-i])
	}
	return c
}

func distance(a, b byte) int {
	if a > b {
		return int(a) - int(b)
	}
	return int(b) - int(a)
}
