package main

import "testing"
import "os"
import "io/ioutil"
import "bytes"

func TestIsCavity(t *testing.T) {
	var cases = []struct {
		m    matrix
		i, j int
		e    bool
	}{
		{matrix{{1, 1, 1, 2}, {1, 9, 1, 2}, {1, 8, 9, 2}, {1, 2, 3, 4}}, 1, 1, true},
		{matrix{{1, 1, 1, 2}, {1, 9, 1, 2}, {1, 8, 9, 2}, {1, 2, 3, 4}}, 2, 1, false},
		{matrix{{1, 1, 1, 2}, {1, 9, 1, 2}, {1, 8, 9, 2}, {1, 2, 3, 4}}, 2, 2, true},
	}
	for _, c := range cases {
		g := c.m.isCavity(c.i, c.j)
		if g != c.e {
			t.Errorf("%s\ni: %d, j: %d, Expected: %t, Got: %t\n", c.m, c.i, c.j, c.e, g)
		}
	}
}

func TestCavityFromFile(t *testing.T) {
	var cases = []struct {
		fis string
		fos string
	}{
		{"input07.txt", "output07.txt"},
	}
	for _, c := range cases {
		fi, _ := os.Open(c.fis)
		fib, _ := ioutil.ReadAll(fi)
		g := readInput(bytes.NewBuffer(fib)).String()
		fo, _ := os.Open(c.fos)
		fob, _ := ioutil.ReadAll(fo)
		e := string(fob)
		if g != e {
			t.Errorf("%s, Expected: %s, Got: %s\n", fib, e, g)
		}
	}
}
