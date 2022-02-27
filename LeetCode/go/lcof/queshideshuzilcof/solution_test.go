package queshideshuzilcof

import "testing"

func TestFunc(t *testing.T) {
	nums := []int{0}
	r := missingNumber(nums)
	if r != 1 {
		t.Fatal(r)
	}
}
