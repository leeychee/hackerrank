package main

import "fmt"

func main() {
	var a0, a1, a2 int
	var b0, b1, b2 int
	fmt.Scanf("%d %d %d\n", &a0, &a1, &a2)
	fmt.Scanf("%d %d %d\n", &b0, &b1, &b2)
	ar0, br0 := compare(a0, b0)
	ar1, br1 := compare(a1, b1)
	ar2, br2 := compare(a2, b2)
	fmt.Printf("%d %d", ar0+ar1+ar2, br0+br1+br2)
}

func compare(i, j int) (ri, rj int) {
	if i == j {
		return 0, 0
	} else if i > j {
		return 1, 0
	} else {
		return 0, 1
	}
}
