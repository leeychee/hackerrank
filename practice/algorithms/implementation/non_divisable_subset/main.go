package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s, k := getCase(os.Stdin)
	fmt.Printf("%d", maxNonDivisableSubsetLen(s, k))
}

type pair struct{ i, j int }

func getCase(r io.Reader) ([]int, int) {
	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &s[i])
	}
	return s, k
}

func maxNonDivisableSubsetLen(s []int, k int) int {
	dp := make(map[int]int)
	pairs := make([]pair, 0, len(s)/2)
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if (s[i]+s[j])%k == 0 {
				pairs = append(pairs, pair{i, j})
				dp[i]++
				dp[j]++
			}
		}
	}
	removed := make(map[int]int)
	for _, p := range pairs {
		i, j := p.i, p.j
		if (s[i]+s[j])%k == 0 {
			if dp[i] > dp[j] {
				removed[i] = dp[i]
			} else {
				removed[j] = dp[j]
			}
		}
	}
	return len(s) - len(removed)
}
