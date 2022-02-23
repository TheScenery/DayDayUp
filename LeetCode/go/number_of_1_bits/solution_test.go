package numberof1bits

import (
	"testing"
)

func TestFunc(t *testing.T) {
	num := uint32(3)
	r := hammingWeight(num)
	if r != 2 {
		t.Fatal(r)
	}
}

func TestFunc1(t *testing.T) {
	num := uint32(3)
	r := hammingWeight1(num)
	if r != 2 {
		t.Fatal(r)
	}
}

func TestFunc2(t *testing.T) {
	num := uint32(3)
	r := hammingWeight2(num)
	if r != 2 {
		t.Fatal(r)
	}
}
