package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var s string
	fmt.Scanf(fmt.Sprintf("%%%ds\n", n), &s)

	fmt.Println(makeBeautiful(s))
}

func makeBeautiful(s string) int {
	var ps []int
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+3] == "010" {
			ps = append(ps, i)
		}
	}
	change := 0
	before := 0
	for i := 0; i < len(ps); i++ {
		if i == 0 || ps[i] != ps[before]+2 {
			change++
			before = i
		}
	}
	return change
}
