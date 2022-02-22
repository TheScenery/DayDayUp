package houserobber

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	r := rob(nums)
	if r != 4 {
		t.Fatal(r)
	}

	nums = []int{2, 7, 9, 3, 1}
	r = rob(nums)
	if r != 12 {
		t.Fatal(r)
	}
}
