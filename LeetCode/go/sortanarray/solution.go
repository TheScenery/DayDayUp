package sortanarray

// 冒泡排序
func sortArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

// 选择排序
func sortArray1(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		minValueIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				minValueIndex = j
			}
		}
		if minValueIndex != i {
			nums[i], nums[minValueIndex] = nums[minValueIndex], nums[i]
		}
	}
	return nums
}

// 快排
func quickSort(nums []int, left, right int) {
	if right <= left {
		return
	}
	key := nums[left]
	i, j := left, right
	for i < j {
		for j > i && nums[j] >= key {
			j--
		}
		for i < j && nums[i] <= key {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	quickSort(nums, i+1, right)
	quickSort(nums, left, i-1)
}

func sortArray2(nums []int) []int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	return nums
}
