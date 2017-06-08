package main

import "testing"

func TestCipher(t *testing.T) {
	var cases = []struct {
		s string
		k int
		e string
	}{
		{"middle-Outz", 2, "okffng-Qwvb"},
	}
	for _, c := range cases {
		g := cipher(c.s, c.k)
		if g != c.e {
			t.Errorf("cipher(%q, %d), Expected: %s, Got: %s", c.s, c.k, c.e, g)
		}
	}
}

func TestCipherB(t *testing.T) {
	var cases = []struct {
		b byte
		k int
		e byte
	}{
		{'a', 1, 'b'},
		{'z', 2, 'b'},
		{'Z', 2, 'B'},
	}
	for _, c := range cases {
		g := cipherB(c.b, c.k)
		if g != c.e {
			t.Errorf("cipherB(%q, %d), Expected: %s, Got: %s", string(c.b), c.k, string(c.e), string(g))
		}
	}
}
