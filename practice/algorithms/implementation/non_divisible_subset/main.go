package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s, k := readInput(os.Stdin)
	m := maxNonDivisibleSubsetLen(s, k)
	fmt.Printf("%d\n", m)
}

func readInput(r io.Reader) ([]int, int) {
	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &s[i])
	}
	return s, k
}

func readOutput(r io.Reader) int {
	var c int
	fmt.Fscanf(r, "%d", &c)
	return c
}

func maxNonDivisibleSubsetLen(s []int, k int) int {
	return maxNonDivisibleSubsetLen3(s, k)
}

func maxNonDivisibleSubsetLen3(s []int, k int) int {
	index := make(map[int][]int)
	for i, v := range s {
		vk := v % k
		if _, ok := index[vk]; !ok {
			index[vk] = make([]int, 0)
		}
		index[vk] = append(index[vk], i)
	}
	m := 0
	used := make(map[int]struct{})
	mid := k/2*2 == k
	for _k, _v := range index {
		if _, ok := used[_k]; ok {
			continue
		}
		if _, ok := used[k-_k]; ok {
			continue
		}
		used[_k] = struct{}{}
		used[k-_k] = struct{}{}
		if _k == 0 || (mid && _k == k/2) {
			m++
		} else {
			js, jok := index[k-_k]
			if jok {
				if len(_v) >= len(js) {
					m += len(_v)
				} else {
					m += len(js)
				}
			} else {
				m += len(_v)
			}
		}
	}
	return m
}

func maxNonDivisibleSubsetLen2(s []int, k int) int {
	m := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		divisible := false
		for j := 0; j < len(m); j++ {
			if (s[i]+m[j])%k == 0 {
				divisible = true
				break
			}
		}
		if !divisible {
			m = append(m, s[i])
		}
	}
	return len(m)
}

type pair struct{ i, j int }

// O(f)=N^2
func maxNonDivisibleSubsetLen1(s []int, k int) int {
	pairs := make([]pair, 0, len(s)/2)
	pairc := make(map[int]int)
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (s[i]+s[j])%k == 0 {
				pairs = append(pairs, pair{i, j})
				pairc[i]++
				pairc[j]++
			}
		}
	}
	rm := make(map[int]int)
	for _, p := range pairs {
		i, j := p.i, p.j
		if pairc[i] > pairc[j] {
			rm[i] = 1
		} else {
			rm[j] = 1
		}
	}
	return len(s) - len(rm)
}
