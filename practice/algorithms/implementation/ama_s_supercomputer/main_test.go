package main

import "testing"
import "bytes"

func TestOverlap(t *testing.T) {
	var cases = []struct {
		p0, p1 *plus
		e      bool
	}{
		{&plus{center: cell{1, 1}}, &plus{center: cell{1, 1}}, true},
		{&plus{center: cell{1, 1}}, &plus{center: cell{2, 1}}, false},
		{&plus{center: cell{1, 2}, d: 1}, &plus{center: cell{2, 1}, d: 1}, true},
		{&plus{center: cell{1, 1}, d: 1}, &plus{center: cell{4, 3}, d: 1}, false},
		{&plus{center: cell{1, 1}, d: 1}, &plus{center: cell{1, 3}, d: 1}, true},
		{&plus{center: cell{1, 1}, d: 1}, &plus{center: cell{2, 3}, d: 1}, false},
		{&plus{center: cell{2, 2}, d: 2}, &plus{center: cell{3, 3}, d: 2}, true},
	}
	for _, c := range cases {
		g := c.p0.overlap(c.p1)
		if g != c.e {
			t.Errorf("%s, %s, Expected: %t, Got: %t", c.p0, c.p1, c.e, g)
		}
	}
}

func TestMaxGoodArea(t *testing.T) {
	var cases = []struct {
		s string
		e int
	}{
		{`5 6
GGGGGG
GBBBGB
GGGGGG
GGBBGB
GGGGGG`, 5},
		{`6 6
BGBBGB
GGGGGG
BGBBGB
GGGGGG
BGBBGB
BGBBGB`, 25},
		{`12 12
GGGGGGGGGGGG
GBGGBBBBBBBG
GBGGBBBBBBBG
GGGGGGGGGGGG
GGGGGGGGGGGG
GGGGGGGGGGGG
GGGGGGGGGGGG
GBGGBBBBBBBG
GBGGBBBBBBBG
GBGGBBBBBBBG
GGGGGGGGGGGG
GBGGBBBBBBBG`, 81},
	}
	for _, c := range cases {
		grid := readInput(bytes.NewReader([]byte(c.s)))
		g := grid.maxGoodArea()
		if g != c.e {
			t.Errorf("%s\n, Expected: %d, Got: %d", c.s, c.e, g)
		}
	}

}
