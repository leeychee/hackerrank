package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	if palindromable(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// If len(s)%2 == 1, one kind of char can be odd;
// Else all kinds of char should be even.
func palindromable(s string) bool {
	bs := []byte(s)
	m := make(map[byte]int)
	for _, b := range bs {
		m[b]++
	}
	center := len(bs) % 2
	for _, v := range m {
		if v%2 != 0 {
			if center == 0 {
				return false
			}
			center--
		}
	}
	return true
}
