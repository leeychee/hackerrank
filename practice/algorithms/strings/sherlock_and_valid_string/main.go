package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	if valid(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func valid(s string) bool {
	bm := make(map[byte]int)
	for _, b := range []byte(s) {
		bm[b]++
	}
	vm := make(map[int]int)
	for _, c := range bm {
		vm[c]++
	}
	if len(vm) == 1 {
		return true
	} else if len(vm) == 2 {
		var v0, v1 int
		for v, k := range vm {
			if k == 1 {
				v0 = v
			} else {
				v1 = v
			}
		}
		return v0 == 1 || (v0 != 0 && (v1-v0 == 1 || v0-v1 == 1))
	} else {
		return false
	}
}
