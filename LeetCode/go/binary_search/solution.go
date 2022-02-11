package binarysearch

func search(nums []int, target int) int {
	length := len(nums)
	start := 0
	end := length - 1
	targetIndex := -1
	for {
		if start > end || targetIndex != -1 {
			break
		}
		testIndex := (end-start)/2 + start
		testValue := nums[testIndex]
		if testValue < target {
			start = testIndex + 1
		} else if testValue > target {
			end = testIndex - 1
		} else {
			targetIndex = testIndex
		}
	}
	return targetIndex
}
