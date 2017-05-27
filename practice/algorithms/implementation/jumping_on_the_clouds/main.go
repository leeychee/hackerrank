package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	s := readInput(os.Stdin)
	if m, err := minJumping(s); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%d\n", m)
	}
}

func readInput(r io.Reader) []int {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	return s
}

func minJumping(s []int) (int, error) {
	r := 0
	for i := 0; i < len(s)-1; {
		if i+2 < len(s) && s[i+2] == 0 {
			i += 2
			r++
		} else if i+1 < len(s) && s[i+1] == 0 {
			i++
			r++
		} else {
			return 0, errors.New("Unable to jumping")
		}
	}
	return r, nil
}
