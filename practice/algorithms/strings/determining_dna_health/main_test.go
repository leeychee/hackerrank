package main

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {

}

func TestHealth(t *testing.T) {
	var cases = []struct {
		gs []gene
		d  dna
		e  int
	}{
		{
			[]gene{{0, "a", 1}, {1, "b", 2}, {2, "c", 3}, {3, "aa", 4}, {4, "d", 5}, {5, "b", 6}},
			"xyz",
			0,
		},
		{
			[]gene{{0, "b", 2}, {1, "c", 3}, {2, "aa", 4}, {3, "d", 5}, {4, "b", 6}},
			"caaab",
			19,
		},
	}
	for _, c := range cases {
		g := newHub(c.gs).health(c.d)
		if g != c.e {
			t.Errorf("%v, %s, Expected: %d, Got: %d", c.gs, c.d, c.e, g)
		}
	}
}

// func BenchmarkHealthFromFile02(b *testing.B) {
// 	benchmarkHealthFromFile(b, "input02.txt")
// }

func BenchmarkHealthFromFile13(b *testing.B) {
	benchmarkHealthFromFile(b, "input13.txt")
}

func BenchmarkHealthFromFile31(b *testing.B) {
	benchmarkHealthFromFile(b, "input31.txt")
}

func benchmarkHealthFromFile(b *testing.B, fn string) {
	f, _ := os.Open(fn)
	genes := readGenes(f)
	cases := readCases(f)
	hub := newHub(genes)
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			hub.sub(c.start, c.end).health(c.dna)
		}
	}
}
