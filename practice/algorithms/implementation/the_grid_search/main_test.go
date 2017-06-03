package main

import "testing"
import "bytes"

func TestGridContain(t *testing.T) {
	var cases = []struct {
		s string
		e error
	}{
		{`4 6
123412
561212
123634
781288
2 2
12
34`, nil},
	}
	for _, c := range cases {
		c0 := readCase(bytes.NewBufferString(c.s))
		_, _, g := c0.g.contain(c0.p)
		if g != c.e {
			t.Errorf("%s, Expected: %v, Got: %v\n", c.s, c.e, g)
		}
	}
}
