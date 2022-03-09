package juzhenzhongdelujinglcof

import "testing"

func TestFunc(t *testing.T) {
	borad := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	r := exist(borad, word)
	if r != true {
		t.Fatal(r)
	}

	borad = [][]byte{{'a', 'b'}, {'c', 'd'}}
	word = "abcd"
	r = exist(borad, word)
	if r != false {
		t.Fatal(r)
	}
}
