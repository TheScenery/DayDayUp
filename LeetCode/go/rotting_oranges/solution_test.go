package rottingoranges

import "testing"

func TestFunc(t *testing.T) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	r := orangesRotting(grid)
	if r != 4 {
		t.Fatal(r)
	}
	grid = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	r = orangesRotting(grid)
	if r != -1 {
		t.Fatal(r)
	}

	grid = [][]int{{0, 2}}
	r = orangesRotting(grid)
	if r != 0 {
		t.Fatal(r)
	}

	grid = [][]int{{0}}
	r = orangesRotting(grid)
	if r != 0 {
		t.Fatal(r)
	}
}
