package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(countWords(s))
}

func countWords(s string) int {
	c := 0
	for i, b := range []byte(s) {
		if i == 0 || (b >= 'A' && b <= 'Z') {
			c++
		}
	}
	return c
}
