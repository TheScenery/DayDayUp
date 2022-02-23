package minimumdifferencebetweenhighestandlowestofkscores

import "testing"

func TestMinDifference(t *testing.T) {
	nums := []int{90}
	k := 1
	r := minimumDifference(nums, k)
	if r != 0 {
		t.Fatal(r)
	}

	nums = []int{9, 4, 1, 7}
	k = 2
	r = minimumDifference(nums, k)
	if r != 2 {
		t.Fatal(r)
	}
}
