package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, _, t := readInput(os.Stdin)
	topics, teams := maxTopicsOfTwoPersons(t)
	fmt.Printf("%d\n%d\n", topics, teams)
}

func readInput(r io.Reader) (int, int, [][]int) {
	var n, m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(r, "%1d", &t[i][j])
		}
		fmt.Fscanf(r, "\n")
	}
	return n, m, t
}

func maxTopicsOfTwoPersons(t [][]int) (topics, teams int) {
	for i := 0; i < len(t)-1; i++ {
		for j := i + 1; j < len(t); j++ {
			ct, _ := or(t[i], t[j])
			if topics < ct {
				topics = ct
				teams = 1
			} else if topics == ct {
				teams++
			}
		}
	}
	return
}

func or(t1, t2 []int) (topics int, t []int) {
	t = make([]int, len(t1))
	for i := 0; i < len(t); i++ {
		if t1[i] == 1 || t2[i] == 1 {
			t[i] = 1
			topics++
		} else {
			t[i] = 0
		}
	}
	return
}
