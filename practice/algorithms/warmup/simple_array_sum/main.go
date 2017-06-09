package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	fmt.Print(simpleSum(s))
}

func simpleSum(s []int) (r int) {
	for _, x := range s {
		r += x
	}
	return
}
