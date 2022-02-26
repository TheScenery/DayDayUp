package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		k := n - 1
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[k]+nums[j] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return result
}
