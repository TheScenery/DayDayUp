package sumofsubarrayranges

func subArrayRanges(nums []int) int64 {
	result := int64(0)
	n := len(nums)
	for i := 0; i < n; i++ {
		min := nums[i]
		max := nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}
			result = result + int64((max - min))
		}
	}
	return result
}
