package zaipaixushuzuzhongchazhaoshuzilcof

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	r := search(nums, target)
	if r != 2 {
		t.Fatal(r)
	}

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	r = search(nums, target)
	if r != 0 {
		t.Fatal(r)
	}

	nums = []int{1}
	target = 1
	r = search(nums, target)
	if r != 1 {
		t.Fatal(r)
	}

	nums = []int{2, 2}
	target = 2
	r = search(nums, target)
	if r != 2 {
		t.Fatal(r)
	}
}
