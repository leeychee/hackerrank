package main

import "testing"
import "reflect"
import "os"

func TestWalk(t *testing.T) {
	var cases = []struct {
		t *trail
		e [][]int
	}{
		{&trail{3, 1, 2}, [][]int{{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {0, 2, 4}}},
		{&trail{4, 10, 100}, [][]int{
			{0, 10, 20, 30},
			{0, 10, 20, 120},
			{0, 10, 110, 120},
			{0, 10, 110, 210},
			{0, 100, 110, 120},
			{0, 100, 110, 210},
			{0, 100, 200, 210},
			{0, 100, 200, 300},
		}},
	}
	for _, c := range cases {
		w := c.t.walk()
		g := make([][]int, 0)
		for {
			g0, ok := <-w
			if !ok {
				break
			}
			g = append(g, g0)
		}
		if !reflect.DeepEqual(g, c.e) {
			t.Errorf("Expected: %v, Got: %v\n", c.e, g)
		}
	}
}

func TestLastsFromFile(t *testing.T) {
	var cases = []struct {
		fis, fos string
	}{
		{"input00.txt", "output00.txt"},
		{"input04.txt", "output04.txt"},
	}
	for _, c := range cases {
		fi, _ := os.Open(c.fis)
		fo, _ := os.Open(c.fos)
		trails := readInput(fi)
		output := readOutput(fo)
		for i := 0; i < len(trails); i++ {
			tc := trails[i]
			g := trails[i].lasts()
			e := output[i]
			if !reflect.DeepEqual(e, g) {
				t.Errorf("(%d, %d, %d), Expected: %v, Got: %v\n", tc.n, tc.a, tc.b, e, g)
			}
		}
	}
}
