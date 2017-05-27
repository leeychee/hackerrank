package main

import "testing"

import "os"

var cases = []struct {
	s []int
	k int
	e int
}{
	{[]int{1, 7, 2, 4}, 3, 3},
}

func TestMaxNonDivisibleSubsetLen(t *testing.T) {
	for _, c := range cases {
		s, k := c.s, c.k
		g := maxNonDivisibleSubsetLen(s, k)
		if g != c.e {
			t.Errorf("Expected: %d, Got: %d", c.e, g)
		}
	}
}

func TestMaxNonDivisibleSubsetLenFromFile(t *testing.T) {
	f, err := os.Open("./input09.txt")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	s, k := readInput(f)
	// t.Log(s)
	of, _ := os.Open("./output09.txt")
	e := readOutput(of)
	g := maxNonDivisibleSubsetLen(s, k)
	if g != e {
		t.Errorf("Expected: %d, Got: %d", e, g)
	}
}

func BenchmarkMaxNonDivisibleSubsetLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxNonDivisibleSubsetLen(cases[0].s, cases[0].k)
	}
}

func BenchmarkMaxNonDivisibleSubsetLenFromFile(b *testing.B) {
	f, _ := os.Open("./input09.txt")
	s, k := readInput(f)
	for i := 0; i < b.N; i++ {
		maxNonDivisibleSubsetLen(s, k)
	}
}
