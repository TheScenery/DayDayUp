package threesum

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	r := threeSum(nums)
	fmt.Println(r)

	nums = []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	r = threeSum(nums)
	fmt.Println(r)
}
