package main

import "testing"
import "math/big"

func TestMFib(t *testing.T) {
	var cases = []struct {
		t0, t1, n int
		e         string
	}{
		{0, 1, 5, "5"},
		{0, 1, 10, "84266613096281243382112"},
	}
	for _, c := range cases {
		g := mFib(c.t0, c.t1, c.n)
		e, _ := big.NewInt(0).SetString(c.e, 10)
		if g.Cmp(e) != 0 {
			t.Errorf("t0: %d, t1: %d, n: %d, Expected: %d, Got: %d", c.t0, c.t1, c.n, e, g)
		}
	}
}
