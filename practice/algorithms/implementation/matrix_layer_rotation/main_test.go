package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadMatrix(t *testing.T) {
	var cases = []struct {
		m, n int
		s    string
	}{
		{4, 4, `1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16`},
	}
	for _, c := range cases {
		m := readMatrix(bytes.NewReader([]byte(c.s)), c.m, c.n)
		if m.String() != c.s {
			t.Errorf("Expected:\n%s\nGot:\n%s", c.s, m.String())
		}
	}
}

func TestMatrixRotate(t *testing.T) {
	var cases = []struct {
		m, n int
		s    string
		e    string
	}{
		{4, 4, `1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16`, `2 3 4 8
1 7 11 12
5 6 10 16
9 13 14 15`},
	}
	for _, c := range cases {
		m := readMatrix(bytes.NewReader([]byte(c.s)), c.m, c.n)
		// fmt.Println(m)
		m.rotate(1)
		if m.String() != c.e {
			t.Errorf("Rotate\n%s\nExpected:\n%s\nGot:\n%s", c.s, c.e, m)
		}
	}
}

func TestRotateFromFile(t *testing.T) {
	var cases = []struct {
		fis, fos string
	}{
		{"input.txt", "output.txt"},
	}
	for _, c := range cases {
		fi, _ := os.Open(c.fis)
		m, r := readInput(fi)
		fo, _ := os.Open(c.fos)
		e, _ := ioutil.ReadAll(fo)
		m.rotate(r)
		if m.String() != string(e) {
			t.Errorf("Expected:\n%s\nGot:\n%s", e, m)
		}
	}
}
