package twosum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
	if result[0] != 1 || result[1] != 0 {
		t.Fail()
	}
}
