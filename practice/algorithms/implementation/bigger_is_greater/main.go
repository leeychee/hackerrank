package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

// ErrNoAnswer no answer err
var ErrNoAnswer = errors.New("no answer")

func main() {
	ss := readInput(os.Stdin)
	for _, s := range ss {
		if t, err := bigger(s); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(t)
		}
	}
}

func readInput(r io.Reader) []string {
	var n int
	fmt.Fscanf(r, "%d\n", &n)

	ss := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%s\n", &ss[i])
	}
	return ss
}

func bigger(s string) (string, error) {
	// when s has only one or zero char, there is no bigger string.
	if len(s) < 2 {
		return "", ErrNoAnswer
	}
	bs := []byte(s)
	n := len(bs)
	rs := []byte{}
	for i := n - 2; i >= 0; i-- {
		if bs[i] >= bs[i+1] {
			if i < 1 {
				return "", ErrNoAnswer
			}
		} else {
			rs = append(rs, bs[:i]...)
			mi, mv := minest(bs[i], bs[i+1:])
			rs = append(rs, mv)
			tail := append([]byte{}, bs[i:i+1+mi]...)
			tail = append(tail, bs[i+2+mi:]...)
			tails := minner(append([]byte{}, tail...))
			rs = append(rs, tails...)
			break
		}
	}
	return string(rs), nil
}

func minest(c byte, bs []byte) (int, byte) {
	v := bs[0]
	r := 0
	for i, b := range bs {
		if b < v && b > c {
			r = i
			v = b
		}
	}
	return r, v
}

func minner(b []byte) []byte {
	is := make([]int, len(b))
	for i, v := range b {
		is[i] = int(v)
	}
	sort.Ints(is)
	r := make([]byte, len(b))
	for i, v := range is {
		r[i] = byte(v)
	}
	return r
}
