package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	bs := readInput(os.Stdin)
	if e, err := makeEven(bs); err != nil {
		fmt.Println("NO")
	} else {
		fmt.Println(e)
	}
}

func readInput(r io.Reader) []int {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	bs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &bs[i])
	}
	return bs
}

// ErrUnableMakeEven unable make even
var ErrUnableMakeEven = errors.New("unable make even")

func makeEven(bs []int) (int, error) {
	e := 0
	for i := 0; i < len(bs)-1; i++ {
		if bs[i]%2 == 1 {
			bs[i]++
			bs[i+1]++
			e += 2
		}
	}
	if bs[len(bs)-1]%2 == 1 {
		return 0, ErrUnableMakeEven
	}
	return e, nil
}
