package main

import "testing"
import "os"
import "io"

var cases = []struct {
	s []int
	k int
	e int
}{
	{[]int{1, 7, 2, 4}, 3, 3},
}

func TestMaxNonDivisableSubsetLen(t *testing.T) {
	for _, c := range cases {
		g := maxNonDivisableSubsetLen(c.s, c.k)
		if g != c.e {
			t.Errorf("Excepted: %q, Got: %q", c.e, g)
		}
	}
}

func BenchmarkMaxNonDivisableSubsetLen(b *testing.B) {
	_if, _ := os.Open("./input09.txt")
	// _of, _ := os.Open("./output09.txt")
	s, k := getCase(io.Reader(_if))
	for i := 0; i < b.N; i++ {
		maxNonDivisableSubsetLen(s, k)
	}
}
