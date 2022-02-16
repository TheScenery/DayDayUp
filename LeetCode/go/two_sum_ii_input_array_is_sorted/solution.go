package twosuminputarrayissorted

func searchIndex(nums []int, target, start, end int) int {
	for {
		if start > end {
			break
		}
		mid := start + (end-start)/2
		midValue := nums[mid]
		if midValue < target {
			start = mid + 1
		} else if midValue > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		y := target - numbers[i]
		yIndex := searchIndex(numbers, y, i+1, n-1)
		if yIndex != -1 {
			return []int{i + 1, yIndex + 1}
		}
	}
	return nil
}
