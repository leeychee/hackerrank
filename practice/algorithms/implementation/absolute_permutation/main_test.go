package main

import "testing"
import "reflect"

var cases = []struct {
	x    *p
	eErr error
	e    []int
}{
	{&p{2, 0}, nil, []int{1, 2}},
	{&p{2, 1}, nil, []int{2, 1}},
	{&p{10, 3}, errNoAbsolute, nil},
	{&p{10, 1}, nil, []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9}},
}

func TestAbsolute(t *testing.T) {
	for _, c := range cases {
		g, gErr := c.x.absolute()
		if gErr != c.eErr || !reflect.DeepEqual(c.e, g) {
			t.Errorf("%s, Expected: %v, Got: %v\n", c.x, c.e, g)
		}
	}
}

func BenchmarkAbsolute10(b *testing.B) {
	p := &p{10, 1}
	for i := 0; i < b.N; i++ {
		p.absolute()
	}
}

func BenchmarkAbsolute1000(b *testing.B) {
	p := &p{1000, 1}
	for i := 0; i < b.N; i++ {
		p.absolute()
	}
}

func BenchmarkAbsolute100000(b *testing.B) {
	p := &p{99112, 2}
	for i := 0; i < b.N; i++ {
		p.absolute()
	}
}
