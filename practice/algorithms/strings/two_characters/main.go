package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var s string
	fmt.Scanf(fmt.Sprintf("%%%ds", n), &s)
	fmt.Println(maxTwoCharacters([]byte(s)))
}

type count struct {
	b byte
	c int
}

type countSorter []count

func (s countSorter) Len() int {
	return len(s)
}

func (s countSorter) Less(i, j int) bool {
	return s[i].c < s[j].c
}

func (s countSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func maxTwoCharacters(bs []byte) int {
	m := make(map[byte]int)
	for i := 0; i < len(bs); i++ {
		m[bs[i]]++
	}
	var counts []count
	for k, v := range m {
		counts = append(counts, count{k, v})
	}
	sort.Sort(countSorter(counts))
	max := 0
	for i := len(counts) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if counts[i].c-counts[j].c < 2 {
				if valid(bs, counts[i].b, counts[j].b) {
					return counts[i].c + counts[j].c
				}
			}
		}
	}
	return max
}

func valid(bs []byte, a, b byte) bool {
	var c byte
	for i := 0; i < len(bs); i++ {
		if bs[i] != a && bs[i] != b {
			continue
		}
		if bs[i] == a && c == a {
			return false
		}
		if bs[i] == b && c == b {
			return false
		}
		c = bs[i]
	}
	return true
}
