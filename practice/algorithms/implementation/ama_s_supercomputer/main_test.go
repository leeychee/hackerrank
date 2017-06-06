package main

import "testing"
import "bytes"
import "reflect"

func TestGridPluses(t *testing.T) {
	var cases = []struct {
		s string
		e []plus
	}{
		{`5 6
GGGGGG
GBBBGB
GGGGGG
GGBBGB
GGGGGG`, []plus{{{2, 1}}, {{1, 4}, {2, 3}, {2, 4}, {2, 5}, {3, 4}}}},
		{`6 6
BGBBGB
GGGGGG
BGBBGB
GGGGGG
BGBBGB
BGBBGB`, []plus{{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}, {{2, 4}, {3, 3}, {3, 4}, {3, 5}, {4, 4}}}},
	}
	for _, c := range cases {
		g := readInput(bytes.NewReader([]byte(c.s)))
		p := g.validPluses()
		if !reflect.DeepEqual(p, c.e) {
			t.Errorf("Expected: %v, Got: %v", c.e, p)
		}
	}
}
