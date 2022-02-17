package knightprobabilityinchessboard

import "testing"

func TestFunc(t *testing.T) {
	r := knightProbability(3, 2, 0, 0)
	if r != 0.0625 {
		t.Fatal(r)
	}

	r = knightProbability(1, 0, 0, 0)
	if r != 1 {
		t.Fatal(r)
	}

	r = knightProbability(8, 30, 6, 4)
	if r != 0.00019052566298333648 {
		t.Fatal(r)
	}
}
