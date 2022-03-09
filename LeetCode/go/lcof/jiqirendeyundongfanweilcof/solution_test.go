package jiqirendeyundongfanweilcof

import "testing"

func TestFunc(t *testing.T) {
	r := movingCount(2, 3, 1)
	if r != 3 {
		t.Fatal(r)
	}

	r = movingCount(3, 1, 0)
	if r != 1 {
		t.Fatal(r)
	}
}
