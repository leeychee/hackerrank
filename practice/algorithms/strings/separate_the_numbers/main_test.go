package main

import (
	"reflect"
	"testing"
)

func TestBeautiful(t *testing.T) {
	var cases = []struct {
		s string
		b []int
		e error
	}{
		{"1234", []int{1, 2, 3, 4}, nil},
		{"0102", nil, errNotBeautiful},
		{"91011", []int{9, 10, 11}, nil},
		{"99100", []int{99, 100}, nil},
		{"9991000", []int{999, 1000}, nil},
		{"99910000", nil, errNotBeautiful},
		{"989910010110", nil, errNotBeautiful},
		{"101103", nil, errNotBeautiful},
		{"010203", nil, errNotBeautiful},
		{"13", nil, errNotBeautiful},
		{"1", nil, errNotBeautiful},
	}
	for _, c := range cases {
		gb, ge := beautiful(c.s)
		if ge != c.e || !reflect.DeepEqual(gb, c.b) {
			t.Errorf("%q, Expected: %v, Got: %v", c.s, c.b, gb)
		}
	}
}
