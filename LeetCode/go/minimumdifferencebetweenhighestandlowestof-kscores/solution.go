package minimumdifferencebetweenhighestandlowestofkscores

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	min := nums[k-1] - nums[0]
	numsLength := len(nums)
	for i := 1; i <= numsLength-k; i++ {
		difference := nums[i+k-1] - nums[i]
		if difference < min {
			min = difference
		}
	}
	return min
}
