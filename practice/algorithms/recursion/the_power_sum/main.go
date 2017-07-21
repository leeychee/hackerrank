package main

import (
	"fmt"
)

func main() {
	var x, n int
	fmt.Scanf("%d\n%d\n", &x, &n)
	fmt.Println(len(expand(x, n)))
}

func expand(x, n int) [][]int {
	rst := make([][]int, 0)
	c := expand0(x, makeCache(x, n), 1)
	for {
		r0, ok := <-c
		if !ok {
			break
		}
		rst = append(rst, r0)
	}
	return rst
}

func expand0(x0 int, c0 []int, cursor int) <-chan []int {
	// fmt.Printf("%d, %v\n", x0, c0[cursor:])
	c := make(chan []int)
	go func() {
		for i := cursor; i < len(c0); i++ {
			v0 := c0[i]
			if v0 == x0 {
				c <- []int{i}
				continue
			} else if v0 > x0 {
				break
			}
			e := expand0(x0-v0, c0, i+1)
			for {
				e0, ok := <-e
				if !ok {
					break
				}
				c <- append([]int{i}, e0...)
			}
		}
		close(c)
	}()
	return c
}

func makeCache(x, n int) []int {
	cache := make([]int, 0)
	for i := 0; ; i++ {
		x0 := power(i, n)
		if x0 > x {
			break
		}
		cache = append(cache, x0)
		if x0 == x {
			break
		}
	}
	return cache
}

func power(x, e int) (y int) {
	y = x
	for i := 1; i < e; i++ {
		y *= x
	}
	return
}
