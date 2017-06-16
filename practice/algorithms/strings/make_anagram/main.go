package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scanf("%s\n", &s1)
	fmt.Scanf("%s\n", &s2)
	fmt.Println(makeAnagramByRemove(s1, s2))
}

func makeAnagramByRemove(s1, s2 string) int {
	m1 := toMap([]byte(s1))
	m2 := toMap([]byte(s2))
	var r int
	for b, v := range m1 {
		if m2[b]-v > 0 {
			r += m2[b] - v
		} else {
			r -= m2[b] - v
		}
		delete(m2, b)
	}
	for b, v := range m2 {
		if m1[b]-v > 0 {
			r += m1[b] - v
		} else {
			r -= m1[b] - v
		}
	}
	return r
}

func toMap(bs []byte) map[byte]int {
	m := make(map[byte]int)
	for _, b := range bs {
		m[b]++
	}
	return m
}
