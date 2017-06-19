package main

import "testing"
import "strings"
import "reflect"

func TestReadStringsN(t *testing.T) {
	var cases = []struct {
		s string
		e []string
	}{
		{"2\nabba\nabcd", []string{"abba", "abcd"}},
	}
	for _, c := range cases {
		g := readStringsN(strings.NewReader(c.s))
		if !reflect.DeepEqual(g, c.e) {
			t.Errorf("%q, Expected: %v, Got: %v", c.s, c.e, g)
		}
	}

}

func TestAnagramCount(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{"aa", 1},
		{"abba", 4},
		{"abcd", 0},
		{"ifailuhkqq", 3},
		{"hucpoltgty", 2},
		{"ovarjsnrbf", 2},
		{"pvmupwjjjf", 6},
		{"iwwhrlkpek", 3},
	}
	for _, c := range cases {
		g := anagramCount(c.s)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}
}
