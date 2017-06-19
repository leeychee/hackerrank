package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	is, k := readInput(os.Stdin)
	p, err := maxPalindrome(is, k)
	if err != nil {
		fmt.Println("-1")
	} else {
		fmt.Println(palindrome(p))
	}
}

func readInput(r io.Reader) ([]int, int) {
	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)

	is := make([]int, n)
	for i := range is {
		fmt.Fscanf(r, "%01d", &is[i])
	}
	return is, k
}

func parseInts(s string) (is []int) {
	for _, b := range []byte(s) {
		is = append(is, int(b-'0'))
	}
	return
}

var errUnablePalindrome = errors.New("unable palindrome")

func maxPalindrome(is []int, change int) ([]int, error) {
	n := len(is)
	minChange := 0
	for i := 0; i < n/2; i++ {
		if is[i] != is[n-1-i] {
			minChange++
		}
	}
	suplus := change - minChange
	if suplus < 0 {
		return nil, errUnablePalindrome
	}
	for i := 0; i <= n/2; i++ {
		// can change a pair
		if is[i] == is[n-1-i] {
			if is[i] != 9 && suplus > 1 {
				is[i] = 9
				is[n-1-i] = 9
				suplus -= 2
			} else if i == n-1-i && suplus > 0 {
				is[i] = 9
			}
		} else {
			if is[i] != 9 && is[n-1-i] != 9 && suplus > 0 {
				is[i] = 9
				is[n-1-i] = 9
				suplus--
			} else if is[i] > is[n-1-i] {
				is[n-1-i] = is[i]
			} else {
				is[i] = is[n-1-i]
			}
		}
	}
	return is, nil
}

type palindrome []int

func (p palindrome) String() string {
	var buf bytes.Buffer
	for _, b := range p {
		buf.WriteByte('0' + byte(b))
	}
	return buf.String()
}
