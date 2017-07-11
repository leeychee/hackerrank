package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	var s1, s2 string
	fmt.Scanf("%s\n", &s1)
	fmt.Scanf("%s\n", &s2)
	fmt.Println(maxCommonSubtringLen(s1, s2))
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}

func maxCommonSubtringLen(s1, s2 string) int {
	psb := trancform(s1, s2)
	return longestIncreasingSeq(psb)
}

// Transform the given issue to "Find the longest increasing sequence issue."
func trancform(s1, s2 string) [][]int {
	m2 := make(map[byte][]int)
	for i, b := range []byte(s2) {
		m2[b] = append(m2[b], i)
		sort.Ints(m2[b])
	}
	var psb [][]int
	for _, b := range []byte(s1) {
		if len(m2[b]) > 0 {
			psb = append(psb, m2[b])
		}
	}
	return psb
}

// Use dynamic programming
// the origin issue also can solve by dynamic programming
func longestIncreasingSeq(matrix [][]int) int {
	max := 0
	for _, v := range matrix {
		last := v[len(v)-1]
		if max < last {
			max = last
		}
	}
	pl := make([]int, max+1)
	for i := len(matrix) - 1; i >= 0; i-- {
		longestIncreasingSeq0(matrix[i], pl)
	}
	max = 0
	for _, v := range pl {
		if max < v {
			max = v
		}
	}
	return max
}

func longestIncreasingSeq0(col []int, pl []int) {
	pl0 := make(map[int]int)
	for _, v := range col {
		max := 0
		for i := v + 1; i < len(pl); i++ {
			if max < pl[i] {
				max = pl[i]
			}
		}
		pl0[v] = max + 1
	}
	for k, v := range pl0 {
		if v > pl[k] {
			pl[k] = v
		}
	}
}
