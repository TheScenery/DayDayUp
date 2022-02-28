package maximumsubarray

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(nums)
	if r != 6 {
		t.Fatal(r)
	}

	nums = []int{1}
	r = maxSubArray(nums)
	if r != 1 {
		t.Fatal(r)
	}

	nums = []int{5, 4, -1, 7, 8}
	r = maxSubArray(nums)
	if r != 23 {
		t.Fatal(r)
	}
}
