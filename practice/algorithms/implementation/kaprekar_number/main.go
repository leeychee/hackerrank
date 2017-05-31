package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	p, q := readInput(os.Stdin)
	ks, err := kaprekar(p, q)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		for i, k := range ks {
			fmt.Printf("%d", k)
			if i < len(ks)-1 {
				fmt.Print(" ")
			}
		}
	}
}

func readInput(r io.Reader) (p, q int) {
	fmt.Fscanf(r, "%d\n", &p)
	fmt.Fscanf(r, "%d\n", &q)
	return
}

// ErrInvalidRange like the name
var ErrInvalidRange = errors.New("INVALID RANGE")

func kaprekar(p, q int) ([]int, error) {
	r := make([]int, 0)
	for i := p; i <= q; i++ {
		if isKarpreka(i) {
			r = append(r, i)
		}
	}
	if len(r) == 0 {
		return nil, ErrInvalidRange
	}
	return r, nil
}

func isKarpreka(x int) bool {
	sqr := x * x
	lt := lenInt(x)
	p := int(math.Pow10(lt))
	l := sqr / p
	r := sqr - l*p
	return r+l == x
}

func lenInt(x int) (l int) {
	for x > 0 {
		l++
		x = x / 10
	}
	return
}
