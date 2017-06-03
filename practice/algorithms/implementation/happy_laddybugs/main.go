package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	games := readInput(os.Stdin)
	for _, g := range games {
		// fmt.Println(g)
		if g.makeHappy() == nil {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

var errUnableHappy = errors.New("unable happy")

type game []byte

func (g game) makeHappy() error {
	m := make(map[byte]int)
	for _, v := range []byte(g) {
		m[v]++
	}
	if m['_'] < 1 {
		if !g.isHappy() {
			return errUnableHappy
		}
		return nil
	}
	for k, v := range m {
		if k != '_' && v == 1 {
			return errUnableHappy
		}
	}
	return nil
}

func (g game) isHappy() bool {
	happy := true
	var inner = false
	var b byte
	var bc int
	for _, v := range []byte(g) {
		if inner {
			if b == v {
				bc++
			} else {
				if bc < 2 {
					happy = false
					break
				}
				if v != '_' {
					inner = true
					b = v
					bc = 1
				} else {
					inner = false
				}
			}
		} else if v != '_' {
			inner = true
			b = v
			bc = 1
		}
	}
	if inner && bc < 2 {
		happy = false
	}
	return happy
}

func readInput(r io.Reader) []game {
	var n int
	fmt.Fscanf(r, "%d\n", &n)

	gs := make([]game, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscanf(r, "%d\n", &k)
		g := make([]byte, k)
		for j := 0; j < k; j++ {
			fmt.Fscanf(r, "%c", &g[j])
		}
		fmt.Fscanln(r)
		gs[i] = g
	}
	return gs
}
