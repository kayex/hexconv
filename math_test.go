package main

import "testing"

func TestIsHex(t *testing.T) {
	var cases = []struct {
		hex string
		exp bool
	}{
		{"0x0", true},
		{"0x1", true},
		{"123", false},
		{"x1", false},
		{"", false},
	}

	for _, c := range cases {
		act := IsHex(c.hex)

		if act != c.exp {
			t.Errorf("Expected IsHex(%q) to return %t, %t given", c.hex, c.exp, act)
		}
	}
}

func TestHexToDec(t *testing.T) {
	var cases = []struct {
		hex string
		exp string
		err bool
	}{
		{"0x0", "0", false},
		{"0x00", "0", false},
		{"0x01", "1", false},
		{"0x10", "16", false},
		{"0x", "", true},
		{"hello", "", true},
		{"", "", true},
	}

	for _, c := range cases {
		act, err := HexToDec(c.hex)

		if err == nil && c.err {
			t.Errorf("Expected HexToDec(%q) to return error, none given.", c.hex)
			continue
		} else if err != nil && !c.err {
			t.Errorf("Expected HexToDec(%q) to return %q, error given: %v", c.hex, c.exp, err)
			continue
		}

		if act != c.exp {
			t.Errorf("Expected HexToDec(%q) to return %q, %q given", c.hex, c.exp, act)
		}
	}
}

func TestDecToHex(t *testing.T) {
	var cases = []struct {
		dec string
		exp string
		err bool
	}{
		{"0", "0x0", false},
		{"1", "0x1", false},
		{"16", "0x10", false},
		{"hello", "", true},
		{"", "", true},
	}

	for _, c := range cases {
		act, err := DecToHex(c.dec)

		if err == nil && c.err {
			t.Errorf("Expected DecToHex(%q) to return error, none given.", c.dec)
			continue
		} else if err != nil && !c.err {
			t.Errorf("Expected DecToHex(%q) to return %q, error given: %v", c.dec, c.exp, err)
			continue
		}

		if act != c.exp {
			t.Errorf("Expected DecToHex(%q) to return %q, %q given", c.dec, c.exp, act)
		}
	}
}
