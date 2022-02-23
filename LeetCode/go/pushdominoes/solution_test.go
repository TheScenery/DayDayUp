package pushdominoes

import (
	"testing"
)

func TestFunc(t *testing.T) {
	dominoes := "RR.L"
	r := pushDominoes(dominoes)
	if r != "RR.L" {
		t.Fatal(r)
	}

	dominoes = ".L.R...LR..L.."
	r = pushDominoes(dominoes)
	if r != "LL.RR.LLRRLL.." {
		t.Fatal(r)
	}
}
