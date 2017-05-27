package main

import "fmt"

func main() {
	var n, x, r int64
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		r += x
	}
	fmt.Print(r)
}
