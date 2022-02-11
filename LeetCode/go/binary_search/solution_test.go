package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	r := search(nums, target)
	if r != 4 {
		t.Fatal(r)
	}

	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 2
	r = search(nums, target)
	if r != -1 {
		t.Fatal(r)
	}
}
