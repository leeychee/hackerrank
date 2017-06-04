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

func TestAbsolute2(t *testing.T) {
	for _, c := range cases {
		g, gErr := c.x.absolute2()
		if gErr != c.eErr || !reflect.DeepEqual(c.e, g) {
			t.Errorf("%s, Expected: %v, Got: %v\n", c.x, c.e, g)
		}
	}
}

func TestAbsolute3(t *testing.T) {
	for _, c := range cases {
		g, gErr := c.x.absolute3()
		if gErr != c.eErr || !reflect.DeepEqual(c.e, g) {
			t.Errorf("%s, Expected: %v, Got: %v\n", c.x, c.e, g)
		}
	}
}

func BenchmarkAbsolute(b *testing.B) {
	p := &p{100, 1}
	for i := 0; i < b.N; i++ {
		p.absolute()
	}
}

func BenchmarkAbsolute2(b *testing.B) {
	p := &p{100, 1}
	for i := 0; i < b.N; i++ {
		p.absolute2()
	}
}

func BenchmarkAbsolute3(b *testing.B) {
	p := &p{100, 1}
	for i := 0; i < b.N; i++ {
		p.absolute3()
	}
}

func BenchmarkAbsolute3_100000(b *testing.B) {
	p := &p{99112, 2}
	for i := 0; i < b.N; i++ {
		p.absolute3()
	}
}
