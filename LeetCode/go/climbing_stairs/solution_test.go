package climbingstairs

import "testing"

func TestFunc(t *testing.T) {
	n := 2
	r := climbStairs(n)
	if r != 2 {
		t.Fatal(r)
	}

	n = 3
	r = climbStairs(n)
	if r != 3 {
		t.Fatal(r)
	}
}
