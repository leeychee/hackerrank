package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ss := readStringsN(os.Stdin)
	for _, s := range ss {
		fmt.Println(makePalindrome(s))
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

var makePalindrome = makePalindrome2

func makePalindrome1(s string) int {
	if isPalindrome(s) {
		return -1
	}
	for i := 0; i < len(s); i++ {
		s0 := s[:i] + s[i+1:]
		if isPalindrome(s0) {
			return i
		}
	}
	return -2
}

func makePalindrome2(s string) int {
	bs := []byte(s)
	n := len(bs)
	for i := 0; i <= n/2; i++ {
		if bs[i] != bs[n-i-1] {
			if isPalindrome(s[i+1 : n-i]) {
				return i
			}
			return n - i - 1
		}
	}
	return -1
}

func isPalindrome(s string) bool {
	bs := []byte(s)
	n := len(bs)
	for i := 0; i <= n/2; i++ {
		if bs[i] != bs[n-i-1] {
			return false
		}
	}
	return true
}
