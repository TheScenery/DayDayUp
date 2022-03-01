package erweishuzuzhongdechazhaolcof

import "testing"

func TestFunc(t *testing.T) {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5
	r := findNumberIn2DArray(matrix, target)
	if r != true {
		t.Fatal(r)
	}

	target = 20
	r = findNumberIn2DArray(matrix, target)
	if r != false {
		t.Fatal(r)
	}
}
