package main

import "testing"

func TestMatrixInfo(t *testing.T) {
	var cases = []struct {
		l, erow, ecol int
	}{
		{54, 7, 8},
		{8, 3, 3},
		{10, 3, 4},
		{12, 3, 4},
	}
	for _, c := range cases {
		grow, gcol := matrixInfo(c.l)
		if gcol != c.ecol || grow != c.erow {
			t.Errorf("Expected: [%d, %d], Got: [%d, %d]\n", c.erow, c.ecol, grow, gcol)
		}
	}
}

func TestEncryption(t *testing.T) {
	var cases = []struct {
		s, e string
	}{
		{"haveaniceday", "hae and via ecy"},
		{"feedthedog", "fto ehg ee dd"},
		{"chillout", "clu hlt io"},
	}
	for _, c := range cases {
		g := encryption(c.s)
		if g != c.e {
			t.Errorf("Expected: %q, Got: %q", c.e, g)
		}
	}
}
