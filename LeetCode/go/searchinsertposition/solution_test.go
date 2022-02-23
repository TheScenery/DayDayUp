package searchinsertposition

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	r := searchInsert(nums, target)
	if r != 2 {
		t.Fatal(r)
	}

	nums = []int{1, 3, 5, 6}
	target = 2
	r = searchInsert(nums, target)
	if r != 1 {
		t.Fatal(r)
	}

	nums = []int{1, 3, 5, 6}
	target = 7
	r = searchInsert(nums, target)
	if r != 4 {
		t.Fatal(r)
	}

	nums = []int{1, 3, 5, 6}
	target = 0
	r = searchInsert(nums, target)
	if r != 0 {
		t.Fatal(r)
	}

	nums = []int{1}
	target = 0
	r = searchInsert(nums, target)
	if r != 0 {
		t.Fatal(r)
	}
}
