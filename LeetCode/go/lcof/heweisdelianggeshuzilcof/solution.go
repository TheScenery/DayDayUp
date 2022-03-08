package heweisdelianggeshuzilcof

func searchIndex(nums []int, target, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func twoSum(nums []int, target int) []int {
	m := len(nums)
	for i, n := range nums {
		v := target - n
		vIndex := searchIndex(nums, v, i+1, m-1)
		if vIndex != -1 {
			return []int{n, nums[vIndex]}
		}
	}
	return nil
}
