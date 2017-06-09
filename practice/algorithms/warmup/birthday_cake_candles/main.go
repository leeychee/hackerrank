package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	s := make([]int, n)
	for i := range s {
		fmt.Scanf("%d", &s[i])
	}
	fmt.Println(maxNumCount(s))
}

func maxNumCount(s []int) int {
	if len(s) < 1 {
		return 0
	}
	max := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
			count = 1
		} else if s[i] == max {
			count++
		}
	}
	return count
}
