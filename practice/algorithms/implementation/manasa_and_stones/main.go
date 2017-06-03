package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	trails := readInput(os.Stdin)
	for _, t := range trails {
		lasts := t.lasts()
		for i, v := range lasts {
			fmt.Print(v)
			if i < len(lasts)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func readInput(r io.Reader) []*trail {
	var t int
	fmt.Fscanf(r, "%d\n", &t)
	ts := make([]*trail, t)
	for i := 0; i < t; i++ {
		t0 := trail{}
		fmt.Fscanln(r, &t0.n)
		fmt.Fscanln(r, &t0.a)
		fmt.Fscanln(r, &t0.b)
		ts[i] = &t0
	}
	return ts
}

func readOutput(r io.Reader) [][]int {
	output := make([][]int, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		lasts := make([]int, 0)
		for _, v := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(v)
			lasts = append(lasts, n)
		}
		output = append(output, lasts)
	}
	return output
}

type trail struct{ n, a, b int }

func (t *trail) lasts() []int {
	a, b := t.a, t.b
	if t.a < t.b {
		a, b = t.b, t.a
	}
	r := make([]int, 0)
	m := make(map[int]struct{})
	for i := 0; i < t.n; i++ {
		r0 := a*i + b*(t.n-1-i)
		if _, ok := m[r0]; !ok {
			r = append(r, r0)
			m[r0] = struct{}{}
		}
	}
	return r
}

func (t *trail) lasts0() []int {
	wC := t.walk()
	lastNumMap := make(map[int]struct{})
	for {
		w, ok := <-wC
		if !ok {
			break
		}
		fmt.Println(w)
		lastNumMap[w[len(w)-1]] = struct{}{}
	}
	keys := make([]int, 0, len(lastNumMap))
	for k := range lastNumMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (t *trail) walk() <-chan []int {
	c0 := make(chan []int)
	go func() {
		seed := []int{0}
		c0 <- seed
		close(c0)
	}()
	c := c0
	for i := 1; i < t.n; i++ {
		c = walk0(c, t.a, t.b)
	}
	return c
}

func walk0(c chan []int, a, b int) chan []int {
	d := make(chan []int)
	go func() {
		for {
			w, ok := <-c
			if !ok {
				close(d)
				break
			}
			w1 := make([]int, len(w)+1)
			copy(w1, w)
			w1[len(w)] = w[len(w)-1] + a
			d <- w1
			w2 := make([]int, len(w)+1)
			copy(w2, w)
			w2[len(w)] = w[len(w)-1] + b
			d <- w2
		}
	}()
	return d
}
