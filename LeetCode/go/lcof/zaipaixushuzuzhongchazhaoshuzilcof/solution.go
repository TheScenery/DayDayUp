package zaipaixushuzuzhongchazhaoshuzilcof

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	mid := (end - start) / 2
	targetStartIndex := -1
	targetEndIndex := -1
	targetIndex := -1
	for start <= end {
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			targetIndex = mid
			break
		}
		mid = start + (end-start)/2
	}
	if targetIndex == -1 {
		return 0
	}
	for i := targetIndex; i >= 0; i-- {
		if nums[i] == target {
			targetStartIndex = i
		} else {
			break
		}
	}
	for i := targetIndex; i < len(nums); i++ {
		if nums[i] == target {
			targetEndIndex = i
		} else {
			break
		}
	}
	return targetEndIndex - targetStartIndex + 1
}
