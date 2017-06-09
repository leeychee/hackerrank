package main

import "testing"

func TestVeryBigSum(t *testing.T) {
	var cases = []struct {
		s []int64
		e string
	}{
		{[]int64{1, 1}, "2"},
		{[]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}, "5000000015"},
		{[]int64{1001458909, 1004570889, 1007019111, 1003302837, 1002514638, 1006431461, 1002575010, 1007514041, 1007548981, 1004402249}, "10047338126"},
	}
	for _, c := range cases {
		g := veryBigSum(c.s).String()
		if g != c.e {
			t.Errorf("%v, Expected: %s, Got: %s", c.s, c.e, g)
		}
	}
}
