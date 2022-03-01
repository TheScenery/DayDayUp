package pattern132

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	r := find132pattern(nums)
	if r != false {
		t.Fatal(r)
	}

	nums = []int{3, 1, 4, 2}
	r = find132pattern(nums)
	if r != true {
		t.Fatal(r)
	}

	nums = []int{-1, 3, 2, 0}
	r = find132pattern(nums)
	if r != true {
		t.Fatal(r)
	}

	nums = []int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10}
	r = find132pattern(nums)
	if r != true {
		t.Fatal(r)
	}
}

func TestFunc1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	r := find132pattern1(nums)
	if r != false {
		t.Fatal(r)
	}

	nums = []int{3, 1, 4, 2}
	r = find132pattern1(nums)
	if r != true {
		t.Fatal(r)
	}

	nums = []int{-1, 3, 2, 0}
	r = find132pattern1(nums)
	if r != true {
		t.Fatal(r)
	}

	nums = []int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10}
	r = find132pattern1(nums)
	if r != true {
		t.Fatal(r)
	}
}
