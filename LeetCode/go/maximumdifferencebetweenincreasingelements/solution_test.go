package maximumdifferencebetweenincreasingelements

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{7, 1, 5, 4}
	r := maximumDifference(nums)
	if r != 4 {
		t.Fatal(r)
	}
}

func TestFunc1(t *testing.T) {
	nums := []int{7, 1, 5, 4}
	r := maximumDifference1(nums)
	if r != 4 {
		t.Fatal(r)
	}
}
