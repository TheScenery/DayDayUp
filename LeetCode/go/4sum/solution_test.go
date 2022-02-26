package foursum

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	r := fourSum(nums, target)
	fmt.Println(r)

	nums = []int{2, 2, 2, 2, 2}
	target = 8
	r = fourSum(nums, target)
	fmt.Println(r)
}
