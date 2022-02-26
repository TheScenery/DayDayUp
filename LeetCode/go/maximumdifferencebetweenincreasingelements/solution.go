package maximumdifferencebetweenincreasingelements

func maximumDifference(nums []int) int {
	maxDifference := -1
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				difference := nums[j] - nums[i]
				if difference > maxDifference {
					maxDifference = difference
				}
			}
		}
	}
	return maxDifference
}

func maximumDifference1(nums []int) int {
	maxDifference := -1
	curMinValue := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] >= curMinValue {
			continue
		}
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				difference := nums[j] - nums[i]
				if difference > maxDifference {
					maxDifference = difference
				}
			}
		}
		curMinValue = nums[i]
	}
	return maxDifference
}
