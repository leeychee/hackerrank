package main

import "testing"

import "strings"
import "reflect"

func TestReadInput(t *testing.T) {
	var cases = []struct {
		s  string
		es []int
		ei int
	}{
		{"4 1\n3943", []int{3, 9, 4, 3}, 1},
		{"6 3\n092282", []int{0, 9, 2, 2, 8, 2}, 3},
	}
	for _, c := range cases {
		gs, gi := readInput(strings.NewReader(c.s))
		if gi != c.ei || !reflect.DeepEqual(gs, c.es) {
			t.Errorf("%q, Expected: %v, %d, Got: %v, %d", c.s, c.es, c.ei, gs, gi)
		}
	}
}

func TestMaxPalindrome(t *testing.T) {
	var cases = []struct {
		is string
		k  int
		ep string
		ee error
	}{
		{"3493", 1, "3993", nil},
		{"0011", 1, "", errUnablePalindrome},
		{"092282", 3, "992299", nil},
		{"12321", 1, "12921", nil},
		{"5", 1, "9", nil},
	}
	for _, c := range cases {
		is := parseInts(c.is)
		ep := parseInts(c.ep)
		gp, ge := maxPalindrome(is, c.k)
		if ge != c.ee || !reflect.DeepEqual(gp, ep) {
			t.Errorf("%s, %d, Expected: %s, %s； Got: %v, %s", c.is, c.k, c.ep, c.ee, palindrome(gp), ge)
		}
	}

}
