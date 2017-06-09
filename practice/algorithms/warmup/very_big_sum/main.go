package main

import "fmt"
import "math/big"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	s := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	fmt.Print(veryBigSum(s))
}

func veryBigSum(s []int64) *big.Int {
	r := big.NewInt(int64(0))
	for _, x := range s {
		r = r.Add(r, big.NewInt(x))
	}
	return r
}
