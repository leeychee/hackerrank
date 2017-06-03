package main

import "testing"

func TestStrangeCounter(t *testing.T) {
	var cases = []struct {
		t, e int
	}{
		{1, 3},
		{3, 1},
		{4, 6},
		{7, 3},
		{14, 8},
		{15, 7},
		{21, 1},
	}
	for _, c := range cases {
		g := counter2(c.t)
		if g != c.e {
			t.Errorf("%d, Expected: %d, Got: %d\n", c.t, c.e, g)
		}
	}
}

func TestCounter(t *testing.T) {
	pc := counter()
	for i := 0; i < 21; i++ {
		p := <-pc
		t.Logf("%s", p)
	}
}

func BenchmarkCounter(b *testing.B) {
	pc := counter()
	for i := 0; i < b.N; i++ {
		<-pc
	}
}

func BenchmarkCounter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter2(i)
	}
}

func BenchmarkCounter3(b *testing.B) {
	pc := counter3()
	for i := 0; i < b.N; i++ {
		<-pc
	}
}

func BenchmarkCounter4(b *testing.B) {
	pc := counter4()
	for i := 0; i < b.N; i++ {
		<-pc
	}
}

func BenchmarkCounter5(b *testing.B) {
	pc := counter5()
	for i := 0; i < b.N; i++ {
		<-pc
	}
}
