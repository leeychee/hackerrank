package main

import "testing"
import "os"
import "io/ioutil"

func TestWeight(t *testing.T) {
	var cases = []struct {
		b byte
		e int
	}{
		{'a', 1},
		{'z', 26},
	}
	for _, c := range cases {
		g := weight(c.b)
		if g != c.e {
			t.Errorf("%q, Expected: %d, Got: %d", string(c.b), c.e, g)
		}
	}
}

func TestUniformWeight(t *testing.T) {
	var cases = []struct {
		fis, fos string
	}{
		{"input02.txt", "output02.txt"},
	}
	for _, c := range cases {
		fi, _ := os.Open(c.fis)
		fo, _ := os.Open(c.fos)
		s, xs := readInput(fi)
		g := toString(s, xs)
		bs, _ := ioutil.ReadAll(fo)
		bs = append(bs, '\n')
		e := string(bs)
		if g != e {
			t.Errorf("Case %s failed.", c.fis)
		}
	}
}
