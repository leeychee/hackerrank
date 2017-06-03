package main

import "testing"

func TestMakeHappy(t *testing.T) {
	var cases = []struct {
		g string
		e error
	}{
		{"RBY_YBR", nil},
		{"X_Y__X", errUnableHappy},
		{"__", nil},
		{"B_RRBR", nil},
		{"MUURDKUE_SBJ__M__XYCRIWD_UMISUY_WUI_W___WBR__YMSSXBXJ_CBB___U_B_UMU__DYXXI_MB__MMD_", errUnableHappy},
		{"VVE", errUnableHappy},
	}
	for _, c := range cases {
		g := game(c.g).makeHappy()
		if g != c.e {
			t.Errorf("%s, Expected: %s, Got: %s\n", c.g, c.e, g)
		}
	}
}

func TestIsHappy(t *testing.T) {
	var cases = []struct {
		g string
		e bool
	}{
		{"RBY_YBR", false},
		{"BBYY", true},
		{"BBY", false},
		{"BB_YY__XXX", true},
		{"", true},
		{"__", true},
		{"_X_", false},
	}
	for _, c := range cases {
		g := game(c.g).isHappy()
		if g != c.e {
			t.Errorf("%s, Expected: %t, Got: %t\n", c.g, c.e, g)
		}
	}
}
