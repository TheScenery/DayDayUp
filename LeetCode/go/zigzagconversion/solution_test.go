package zigzagconversion

import "testing"

func TestFunc(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	r := convert(s, numRows)
	if r != "PAHNAPLSIIGYIR" {
		t.Fatal(r)
	}
	s = "PAYPALISHIRING"
	numRows = 4
	r = convert(s, numRows)
	if r != "PINALSIGYAHRPI" {
		t.Fatal(r)
	}
	s = "A"
	numRows = 1
	r = convert(s, numRows)
	if r != "A" {
		t.Fatal(r)
	}
	s = "AB"
	numRows = 1
	r = convert(s, numRows)
	if r != "AB" {
		t.Fatal(r)
	}
}
