package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	n, C := readInput(os.Stdin)
	fmt.Printf("%d, %v\n", n, C)
	sort.Ints(C)
	c := changes(C, n)
	count := 0
	for {
		_, ok := <-c
		if !ok {
			break
		}
		count++
	}
	fmt.Println(count)
}

func readInput(r io.Reader) (n int, C []int) {
	var m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	C = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(r, "%d", &C[i])
	}
	return
}

func changes(C []int, n int) <-chan []int {
	c := make(chan []int)
	return c
}
