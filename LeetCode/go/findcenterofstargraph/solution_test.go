package findcenterofstargraph

import "testing"

func TestFunc(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	r := findCenter(edges)
	if r != 2 {
		t.Fatal(r)
	}
}
