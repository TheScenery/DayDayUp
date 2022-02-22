package triangle

import "testing"

func TestFunc(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	r := minimumTotal(triangle)
	if r != 11 {
		t.Fatal(r)
	}

	triangle = [][]int{{-10}}
	r = minimumTotal(triangle)
	if r != -10 {
		t.Fatal(r)
	}
}
