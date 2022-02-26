package foursum

import "sort"

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			newTarget := target - nums[i] - nums[j]
			l := n - 1
			for k := j + 1; k < n; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				for k < l && nums[k]+nums[l] > newTarget {
					l--
				}
				if k == l {
					break
				}
				if nums[k]+nums[l] == newTarget {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
				}
			}
		}
	}
	return result
}
