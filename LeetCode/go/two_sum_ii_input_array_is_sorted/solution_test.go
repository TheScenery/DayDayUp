package twosuminputarrayissorted

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	r := twoSum(nums, target)
	fmt.Println(r)

	nums = []int{2, 3, 4}
	target = 6
	r = twoSum(nums, target)
	fmt.Println(r)
}
