package singleelementinasortedarray

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	r := singleNonDuplicate(nums)
	if r != 2 {
		t.Fatal(r)
	}
	nums = []int{3, 3, 7, 7, 10, 11, 11}
	r = singleNonDuplicate(nums)
	if r != 10 {
		t.Fatal(r)
	}
}
