package main

import (
	"fmt"
	"math/big"
)

func main() {
	var t1, t2, n int
	fmt.Scanf("%d %d %d\n", &t1, &t2, &n)
	f := mFib(t1, t2, n)
	fmt.Println(f)
}

var t []*big.Int

func mFib(t1, t2, n int) (f *big.Int) {
	if t == nil {
		t = []*big.Int{nil, big.NewInt(int64(t1)), big.NewInt(int64(t2))}
	}
	if len(t) > n {
		return t[n]
	}
	tn1 := mFib(t1, t2, n-1)
	tn2 := mFib(t1, t2, n-2)
	stn1 := big.NewInt(1).Mul(tn1, tn1)
	f = big.NewInt(0).Add(tn2, stn1)
	t = append(t, f)
	return
}
