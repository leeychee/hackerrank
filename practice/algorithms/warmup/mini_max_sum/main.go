package main

import (
	"fmt"
	"sort"
)

func main() {
	s := make([]int, 5)
	for i := range s {
		fmt.Scanf("%d", &s[i])
	}
	fmt.Printf("%d %d", minSum(s), maxSum(s))
}

func minSum(s []int) (r int) {
	s0 := make([]int, len(s))
	copy(s0, s)
	sort.Ints(s0)
	for i := 0; i < len(s)-1; i++ {
		r += s0[i]
	}
	return
}

func maxSum(s []int) (r int) {
	s0 := make([]int, len(s))
	copy(s0, s)
	sort.Ints(s0)
	for i := 1; i < len(s); i++ {
		r += s0[i]
	}
	return
}
